package main

import "github.com/qingwenjie/logger"

func main() {
	logger.Log = logger.New(&logger.Options{
		Level:       logger.DefaultLevel,
		LogFilePath: "",
	})
	t()
}

func t() {
	logger.Log.Info("Info info")
	logger.Log.Infof("Info info%s-%d","tttt",111)
	logger.Log.Error("Error info")
	logger.Log.Debug("Debug info")
	logger.Log.Trace("Trace info")
	logger.Log.Panic("Panic info")
	logger.Log.Fatal("Fatal info")
}
