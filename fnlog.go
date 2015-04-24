// logger can print function name and file, line info
package fnlog

import (
	"log"
	"path/filepath"
	"runtime"
)

var (
	DEBUG bool = false
)

// Should implement io.Writer interface
type Logger struct{}

// Implementation of io.Writer for our own logger
// Params:
//     p: Input log message
// Return:
//     n: Count of bytes processed
//     err: Error
func (l Logger) Write(p []byte) (n int, err error) {
	pc, file, line, ok := runtime.Caller(4)
	if DEBUG {
		log.Printf("pc: %v, file: %v, line: %v, ok: %v", pc, file, line, ok)
	}

	if !ok {
		log.Printf("flog.Logger.Write(): runtime.Caller(4): error")
		log.Printf("%s", p)
	} else {
		f := filepath.Base(file)
		fn := runtime.FuncForPC(pc)
		fnName := filepath.Base(fn.Name())
		log.Printf("%s:%d %s(): %s", f, line, fnName, p)
	}

	return len(p), nil
}

// New a logger which can print file, function, line
// Params:
//     tag: Tag will be added before log message
// Return:
//     *log.Logger that can call any function(Ex: Printf(), Println()...)
func New(tag string) *log.Logger {
	prefix := ""
	if tag != "" {
		prefix = tag + ": "
	}
	return log.New(Logger{}, prefix, 0)
}
