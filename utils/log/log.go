package log

import (
	"fmt"
	"log"
	"os"
)

// Logger represent common interface for logging function
type Logger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorGray   = "\033[37m"
	ColorWhite  = "\033[37m"
)

type CustomLog struct {
	Output *os.File
}

func NewLogger(output *os.File) Logger {
	return &CustomLog{output}
}

func (l *CustomLog) Info(args ...interface{}) {
	log.New(l.Output, fmt.Sprint(ColorCyan, "[INFO] ", ColorReset), log.Ldate|log.Ltime|log.Lshortfile).Output(2, fmt.Sprintf("%v", args))
}

func (l *CustomLog) Debug(args ...interface{}) {
	log.New(l.Output, fmt.Sprint(ColorGray, "[DEBUG] ", ColorReset), log.Ldate|log.Ltime|log.Lshortfile).Output(2, fmt.Sprintf("%v", args))
}

func (l *CustomLog) Warn(args ...interface{}) {
	log.New(l.Output, fmt.Sprint(ColorYellow, "[WARN] ", ColorReset), log.Ldate|log.Ltime|log.Lshortfile).Output(2, fmt.Sprintf("%v", args))
}

func (l *CustomLog) Error(args ...interface{}) {
	log.New(l.Output, fmt.Sprint(ColorRed, "[ERROR] ", ColorReset), log.Ldate|log.Ltime|log.Lshortfile).Output(2, fmt.Sprintf("%v", args))
}

func (l *CustomLog) Fatal(args ...interface{}) {
	log.New(l.Output, fmt.Sprint(ColorRed, "[FATAL] ", ColorReset), log.Ldate|log.Ltime|log.Lshortfile).Output(2, fmt.Sprintf("%v", args))
}

func (l *CustomLog) Infof(format string, args ...interface{}) {
	log.New(l.Output, fmt.Sprint(ColorCyan, "[INFO] ", ColorReset), log.Ldate|log.Ltime|log.Lshortfile).Output(2, fmt.Sprintf(format, args))
}

func (l *CustomLog) Debugf(format string, args ...interface{}) {
	log.New(l.Output, fmt.Sprint(ColorGray, "[DEBUG] ", ColorReset), log.Ldate|log.Ltime|log.Lshortfile).Output(2, fmt.Sprintf(format, args))
}

func (l *CustomLog) Warnf(format string, args ...interface{}) {
	log.New(l.Output, fmt.Sprint(ColorYellow, "[WARN] ", ColorReset), log.Ldate|log.Ltime|log.Lshortfile).Output(2, fmt.Sprintf(format, args))
}

func (l *CustomLog) Errorf(format string, args ...interface{}) {
	log.New(l.Output, fmt.Sprint(ColorRed, "[ERROR] ", ColorReset), log.Ldate|log.Ltime|log.Lshortfile).Output(2, fmt.Sprintf(format, args))
}

func (l *CustomLog) Fatalf(format string, args ...interface{}) {
	log.New(l.Output, fmt.Sprint(ColorRed, "[FATAL] ", ColorReset), log.Ldate|log.Ltime|log.Lshortfile).Output(2, fmt.Sprintf(format, args))
}
