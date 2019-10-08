package flogging

import (
	"go.uber.org/zap/zapcore"
)

const (
	defaultFormat = "%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}"
	defaultLevel  = zapcore.InfoLevel
)

var Global *Logging

func init() {
	logging, err := New(Config{})
	if err != nil {
		panic(err)
	}

	Global = logging
}

func Init(config Config) {
	err := Global.Apply(config)
	if err != nil {
		panic(err)
	}
}

func MustGetlogger(loggerName string) *FabricLogger {
	return Global.Logger(loggerName)
}
