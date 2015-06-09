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
type logger struct{}

// Write is the implementation of io.Writer for our own logger
func (l logger) Write(p []byte) (n int, err error) {
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

// New create a new logger which can print file, function, line.
// Params:
//     tag: Tag will be added before log message.
// Return:
//     *log.Logger that can call any log.Logger function(Ex: Printf(), Println()...).
func New(tag string) *log.Logger {
	prefix := ""
	if tag != "" {
		prefix = tag + ": "
	}
	return log.New(logger{}, prefix, 0)
}
