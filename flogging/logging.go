package flogging

import (
	"github.com/w862456671/ZapTest/flogging/fabenc"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"sync"
)

type Config struct {
	// Format is the log record format specifier for the Logging instance. If the
	// spec is the string "json", log records will be formatted as JSON. Any
	// other string will be provided to the FormatEncoder. Please see
	// fabenc.ParseFormat for details on the supported verbs.
	//
	// If Format is not provided, a default format that provides basic information will
	// be used.
	Format string

	// LogSpec determines the log levels that are enabled for the logging system. The
	// spec must be in a format that can be processed by ActivateSpec.
	//
	// If LogSpec is not provided, loggers will be enabled at the INFO level.
	LogSpec string

	// Writer is the sink for encoded and formatted log records.
	//
	// If a Writer is not provided, os.Stderr will be used as the log sink.
	Writer io.Writer
}

type Logging struct {
	*LoggerLevels

	mutex          sync.RWMutex
	encoding       Encoding
	encoderConfig  zapcore.EncoderConfig
	multiFormatter *fabenc.MultiFormatter
	writer         zapcore.WriteSyncer
}

func New(c Config) (*Logging, error) {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.NameKey = "name"

	s := &Logging{
		encoderConfig: encoderConfig,
		LoggerLevels: &LoggerLevels{
			defaultLevel: defaultLevel,
		},
		multiFormatter: fabenc.NewMultiFormatter(),
	}

	err := s.Apply(c)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Logging) Apply(c Config) error {
	err := s.SetFormat(c.Format)
	if err != nil {
		return err
	}

	if c.LogSpec == "" {
		c.LogSpec = os.Getenv("FABRIC_LOGGING_SPEC")
	}
	if c.LogSpec == "" {
		c.LogSpec = defaultLevel.String()
	}

	err = s.LoggerLevels.ActivateSpec(c.LogSpec)
	if err != nil {
		return err
	}

	//	if c.Writer == nil {
	//		c.Writer = os.Stderr
	//	}
	s.SetWriter(c.Writer)

	//	var formatter logging.Formatter
	//	if s.Encoding() == JSON {
	//		formatter = SetFormat(defaultFormat)
	//	} else {
	//		formatter = SetFormat(c.Format)
	//	}
	//
	//	InitBackend(formatter, c.Writer)

	return nil
}

func (s *Logging) SetFormat(format string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if format == "" {
		format = defaultFormat
	}

	if format == "json" {
		s.encoding = JSON
		return nil
	}

	formatters, err := fabenc.ParseFormat(format)
	if err != nil {
		return err
	}
	s.multiFormatter.SetFormatters(formatters)
	s.encoding = CONSOLE

	return nil
}

func (s *Logging) SetWriter(w io.Writer) {
	var sw zapcore.WriteSyncer
	switch t := w.(type) {
	case *os.File:
		sw = zapcore.Lock(t)
	case zapcore.WriteSyncer:
		sw = t
	default:
		hook := lumberjack.Logger{
			Filename:   "./test/test.log",
			MaxAge:     7,
			MaxSize:    1,
			MaxBackups: 1,
			Compress:   true,
		}
		sw = zapcore.AddSync(&hook)
	}

	s.mutex.Lock()
	s.writer = sw
	s.mutex.Unlock()
}

// Write satisfies the io.Write contract. It delegates to the writer argument
// of SetWriter or the Writer field of Config. The Core uses this when encoding
// log records.
// Sync satisfies the zapcore.WriteSyncer interface. It is used by the Core to
// flush log records before terminating the process.
func (s *Logging) Write(b []byte) (int, error) {
	s.mutex.RLock()
	w := s.writer
	s.mutex.RUnlock()

	return w.Write(b)
}
func (s *Logging) Sync() error {
	s.mutex.RLock()
	w := s.writer
	s.mutex.RUnlock()

	return w.Sync()
}

// Encoding satisfies the Encoding interface. It determines whether the JSON or
// CONSOLE encoder should be used by the Core when log records are written.
func (s *Logging) Encoding() Encoding {
	s.mutex.RLock()
	e := s.encoding
	s.mutex.RUnlock()
	return e
}

func (s *Logging) ZapLogger(name string) *zap.Logger {
	levelEnabler := zap.LevelEnablerFunc(func(l zapcore.Level) bool { return true })

	s.mutex.RLock()
	core := &Core{
		LevelEnabler: levelEnabler,
		Levels:       s.LoggerLevels,
		Encoders: map[Encoding]zapcore.Encoder{
			JSON:    zapcore.NewJSONEncoder(s.encoderConfig),
			CONSOLE: fabenc.NewFormatEncoder(s.multiFormatter),
		},
		Selector: s,
		Output:   s,
	}
	s.mutex.RUnlock()

	return NewZapLogger(core).Named(name)
}

func (s *Logging) Logger(name string) *FabricLogger {
	zl := s.ZapLogger(name)
	return NewFabricLogger(zl)
}
