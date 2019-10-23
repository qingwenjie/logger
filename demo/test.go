package main

import "utils/logger"

func main() {
	logger.TextLog = logger.New(&logger.Options{
		LogType:     "file",
		Level:       logger.ErrorLevel,
		RuntimePath: "a",
	})
	logger.TextLog.Error("111")
}
