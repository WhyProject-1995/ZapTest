package main

import (
	"fmt"
	//	logg "github.com/hyperledger/fabric/common/flogging"
	"github.com/w862456671/zap_test/flogging"
	"github.com/w862456671/zap_test/flogging/httpadmin"
	"net/http"
	"time"
)

//var test, _ = os.OpenFile("test.log", os.O_RDWR|os.O_CREATE, 0644)
var logger = flogging.MustGetlogger("FileLog")

func initializeLogging() {
	loggingSpec := "DEBUG"
	loggingFormat := ""
	flogging.Init(flogging.Config{
		Format: loggingFormat,
		//	Writer: test,
		LogSpec: loggingSpec,
	})
}

func main() {
	initializeLogging()

	for i := 0; ; i++ {
		time.Sleep(1 * time.Second)
		logger.Info("ed5f71fcae3420b4509b227239865ecff3ffe8b6e8cecdd38408cd1474fea5e"+
			"5ed5f71fcae3420b4509b227239865ecff3ffe8b6e8cecdd38408cd1474fea5e5ed5f71fcae3420b4509"+
			"b227239865ecff3ffe8b6e8cecdd38408cd1474fea5e5ed5f71fcae3420b4509b227239865ecff3ffe8b6e8"+
			"cecdd38408cd1474fea5e5ed5f7ed5f71fcae3420b4509b227239865ecff3ffe8b6e8cecdd38408cd1474fea5e51fcae3420b4509b22"+
			"ed5f71fcae3420b4509b22723986ed5f71fcae3420b4cecdd38408cd1474fea5e5ed5f"+
			"7ed5f71fcae3420b4509b227239865ecff3ffe8b6e8cecdd38408ccecdd38408cd1474f"+
			"ea5e5ed5f7ed5f71fcae3420b4509b227239865ecff3ffe8b6e8cecdd38408ccecdd38408cd14"+
			"74fea5e5ed5f7ed5f71fcae3420b4509b227239865ecff3ffe8b6e8cecdd38408ccecdd38408cd1"+
			"474fea5e5ed5f7ed5f71fcae3420b4509b227239865ecff3ffe8b6e8cecdd38408ccecdd38408cd1"+
			"474fea5e5ed5f7ed5f71fcae3420b4509b227239865ecff3ffe8b6e8cecdd38408ccecdd38408cd14"+
			"74fea5e5ed5f7ed5f71fcae3420b4509b227239865ecff3ffe8b6e8cecdd38408ccecdd38408cd1474"+
			"fea5e5ed5f7ed5f71fcae3420b4509b227239865ecff3ffe8b6e8cecdd38408ccecdd38408cd1474fea"+
			"5e5ed5f7ed5f71fcae3420b4509b227239865ecff3ffe8b6e8cecdd38408ccecdd38408cd1474fea5e5ed"+
			"5f7ed5f71fcae3420b4509b227239865ecff3ffe8b6e8cecdd38408ccecdd38408cd1474fea5e5ed5f7ed"+
			"5f71fcae3420b4509b227239865ecff3ffe8b6e8cecdd38408ccecdd38408cd1474fea5e5ed5f7ed5f71fca"+
			"e3420b4509b227239865ecff3ffe8b6e8cecdd38408ccecdd38408cd1474fea5e5ed5f7ed5f71fcae3420b4509"+
			"b227239865ecff3ffe8b6e8cecdd38408ccecdd38408cd1474fea5e5ed5f7ed5f71fcae3420b4509b227239865ec"+
			"ff3ffe8b6e8cecdd38408ccecdd38408cd1474fea5e5ed5f7ed5f71fcae3420b4509b227239865ecff3ffe8b6e8cecd"+
			"d38408ccecdd38408cd1474fea5e5ed5f7ed5f71fcae3420b4509b227239865ecff3ffe8b6e8cecdd38408c509b2272"+
			"39865ecff3ffe8b6e8cecdd38408cd1474fea5e55ecff3ffe8b6e8cecdd38408cd1"+
			"474fea5e5ed5f71fcae3420b4509b2ed5cecdd38408cd1474fea5e5ed5f7ed5f71fcae3420b4509b227239865ecff3ffe8b6e8cecdd38408c"+
			"f71fcae3420b4509b227239865ecff3ffe8b6e8cecdd38408cd1474fea5e527239865ecff3ffe8b6e8cecdd38408cd"+
			"ed5f71fcae3420b4509b22723ed5f71fcae3420b4509b227239865ecff3ffe8b6e8cecdd38408cd1474fea5e5ed5f71fcae3420b"+
			"4509b227239865ecff3ffe8b6eed5f71fcae3420b4509b227239865ecff3ffe8b6e8cecdd38408cd1474fea5e58cecdd38408cd1474fe"+
			"a5e5ed5f71fcae3420b4509b22ed5f71fcae3420b4509b227239865ecff3ffe8b6e8cecdd38408cd1474fea5e57239865ecff3ffe8b6e8cecdd38408cd1474f"+
			"ea5e5ed5f71fcae3420b4509b227239865ecff3ffe8b6e8cecdd38408cd1474fea5e5ed5f71fcae3420b4509b227239865ecff3ffe8b6e8"+
			"cecdd38408cd1474fea5e5ed5f71fcae3420b4509b227239865ecff3ffe8b6e8cecdd38408cd1474fea5e5ed5f71fcae3420b"+
			"4509b227239865ecff3ffe8b6e8ceed5f71fcae3420b4509b227239865ecff3ffe8b6e8cecdd38408cd1474fea5e5cdd38408cd1474fea5e5ed5f71fcae3420b4"+
			"509b227239865ecff3ffe8b6e8cecdd38408cd1474fea5e5ed5f71fcae3420b4509b227239865ecff3ffe8b6e8c"+
			"ecdd38408cd1474fea5e59865ecff3ffe8b6e8ceced5f71fcae3420b4509b227239865ecff3ffe8b6e8cecdd38408cd1474fea5e5dd38408cd1474fea5e5ed5f71fcae34"+
			"20b4509b227239865ecff3ffe8b6e8cecdd38408cd1474fea5e51474fea5e57239865ecff3ffeed5f71fcae3420b4509b227239865ecff3ffe8b6e8cecdd38408cd1474fea5e5"+
			"8b6e8cecdd38408cd1474fea5e5 test FileLog ", i)
	}
}

func httpserver() {
	sh := httpadmin.NewSpecHandler()
	mux := http.NewServeMux()
	mux.Handle("/logspec", sh)
	err := http.ListenAndServe("192.168.0.155:8080", sh)
	if err != nil {
		fmt.Println(err)
		return
	}
}
