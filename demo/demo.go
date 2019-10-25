package main

import "logger"

func main() {
	logger.TextLog = logger.New(&logger.Options{
		Level:       logger.DefaultLevel,
		LogType:     logger.LOG_TYPE_FILE,
		LogFilePath: "",
	})
	t()
}

func t() {
	logger.TextLog.Info("test info")
}
