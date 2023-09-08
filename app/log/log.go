package log

import "go.uber.org/zap"

var logger *zap.Logger

func InitLog() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)
	defer logger.Sync() // flushes buffer, if any

}
