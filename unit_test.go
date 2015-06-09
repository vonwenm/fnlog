package fnlog_test

import (
	"github.com/northbright/fnlog"
	"log"
)

var (
	noTag *log.Logger
)

func Example() {
	i := fnlog.New("i", true, false, false)
	w := fnlog.New("w", true, false, false)
	e := fnlog.New("e", true, false, true)

	i.Printf("infos")
	w.Printf("warnnings")
	e.Printf("errors")
	noTag.Printf("no tag")

	// Output:
	// i: infos
	// w: warnnings
	// e: unit_test.go:19 fnlog_test.Example(): errors
	// no tag
}

func init() {
	noTag = fnlog.New("", true, false, false)
}
