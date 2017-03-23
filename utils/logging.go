package utils

import (
	"io"
	"log"
)

var (
	Trace    *log.Logger
	Info     *log.Logger
	Question *log.Logger
	Fail     *log.Logger
	Warning  *log.Logger
	Error    *log.Logger
)

//Init the logging function
func Init(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle, "\033[33m[*] \033[0m", 0)
	Info = log.New(infoHandle, "\033[32m[+] \033[0m", 0)
	Fail = log.New(infoHandle, "\033[91m[x] \033[0m", 0)
	Question = log.New(infoHandle, "\033[91m[?] \033[0m", 0)
	Warning = log.New(warningHandle,
		"\033[91m[WARNING] \033[0m", 0)

	Error = log.New(errorHandle,
		"\033[31mERROR\033[0m: ", log.Ldate|log.Ltime)
}
