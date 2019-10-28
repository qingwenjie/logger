package logger

import "testing"

func TestInfo(t *testing.T) {
	Text = New(&Options{
		Level:       DefaultLevel,
		LogType:     "text",
		LogFilePath: "runtime",
	})

	Text.Info("I'm info")
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
