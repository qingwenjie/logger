package logger

import "testing"

func TestInfo(t *testing.T) {
	TextLog = New(&Options{
		Level:       DefaultLevel,
		LogType:     "text",
		LogFilePath: "runtime",
	})

	TextLog.Info("I'm info")
}

//func Test_DefaultOptions(t *testing.T) {
//
//}
//
//func Test_Infof(t *testing.T) {
//
//}
//
//func Test_Error(t *testing.T) {
//
//}
