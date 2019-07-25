package xmpp 

import "fmt"

var lg Logger


func init() {
	lg = &DefaultLogger{}
}

func SetLogger(logger Logger) {
	lg = logger
}

type Level int

const (
	DEBUG Level = iota
	TRACE
	INFO
	WARNING
	ERROR
)

type Logger interface {
	Log(level Level, format string, args ...interface{})
}


type DefaultLogger struct{
	
}

func (l *DefaultLogger) Log(level Level, format string, args ...interface{}) {
	fmt.Printf(format, args)
}



