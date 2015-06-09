package fnlog_test

import (
	"github.com/northbright/fnlog"
	"log"
)

func Example() {
	// New a *log.Logger
	iLog := fnlog.New("i")
	wLog := fnlog.New("w")
	eLog := fnlog.New("e")
	var noTagLog *log.Logger = fnlog.New("")

	iLog.Printf("print infos")
	wLog.Printf("print warnnings")
	eLog.Printf("print errors")
	noTagLog.Printf("print messages without tag")

	// Output:
}
