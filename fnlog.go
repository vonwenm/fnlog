package fnlog

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var (
	DEBUG bool = false
)

// Customized Logger.
//
//   tag - Users' tag at the beginning of log message. Ex: ERR, DEBUG, INFO...
//   redirectToStdout - if redirect from os.Stderr to os.Stdout.
//   showDateTime - if set flags to log.LstdFlags. See log.Flags() for more information.
//   showDetails - if show package name, file name, func name and line number.
type logger struct {
	tag              string
	redirectToStdout bool
	showDateTime     bool
	showDetails      bool
}

// Write is the implementation of io.Writer for our own logger.
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

		// Set Output to os.Stdout if need
		if l.redirectToStdout {
			log.SetOutput(os.Stdout)
		} else {
			log.SetOutput(os.Stderr)
		}

		// Add tag(prefix) if need
		if l.tag != "" {
			prefix := fmt.Sprintf("%s: ", l.tag)
			log.SetPrefix(prefix)
		} else {
			log.SetPrefix("")
		}

		// Reset flags
		flags := 0
		if l.showDateTime {
			flags = log.LstdFlags
		}
		log.SetFlags(flags)

		// Check if show details
		if l.showDetails {
			log.Printf("%s:%d %s(): %s", f, line, fnName, p)
		} else {
			log.Printf("%s", p)
		}
	}

	return len(p), nil
}

// New create a new log.Logger which can print users' tag package / file / function name and line number.
//
//   Params:
//       tag -  Users' tag at the beginning of log message. Ex: ERR, DEBUG, INFO...
//       redirectToStdout - if redirect from os.Stderr to os.Stdout.
//       showDateTime - if show date / time prefix.
//       showDetails - if show package name, file name, func name and line number.
//   Return:
//       log.Logger that can call any log.Logger function(Ex: Printf(), Println()...).
func New(tag string, redirectToStdout, showDateTime, showDetails bool) *log.Logger {
	l := logger{tag, redirectToStdout, showDateTime, showDetails}
	return log.New(l, "", 0)
}
