package log

import "os"

var globalLog = NewLogger(os.Stdout)

func Info(args ...any) {
	globalLog.Info(args)
}

func Error(args ...any) {
	globalLog.Error(args)
}
