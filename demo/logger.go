package main

import "github.com/qingwenjie/logger"

func main() {
	logger.Text = logger.New(&logger.Options{
		Level:       logger.DefaultLevel,
		LogType:     logger.LOG_TYPE_FILE,
		LogFilePath: "",
	})
	t()
}

func t() {
	logger.Text.Info("test info")
}
