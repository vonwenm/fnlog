package fnlog

import (
	"log"
	"testing"
)

func TestLogger(t *testing.T) {
	log.Printf("Testing fnlog.New())...\n")

	iLog := New("i")
	wLog := New("w")
	eLog := New("e")
	noTagLog := New("")

	iLog.Printf("print infos")
	wLog.Printf("print warnnings")
	eLog.Printf("print errors")
	noTagLog.Printf("print messages without tag")
}
