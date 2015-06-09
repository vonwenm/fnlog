package fnlog_test

import (
	"github.com/northbright/fnlog"
	"log"
)

/*var (
	noTagLog *log.Logger
)
*/
func Example() {
	iLog := fnlog.New("i")
	wLog := fnlog.New("w")
	eLog := fnlog.New("e")

	iLog.Printf("print infos")
	wLog.Printf("print warnnings")
	eLog.Printf("print errors")
	//noTagLog.Printf("print messages without tag")

	// Output:
	// 2015/06/09 14:32:59 unit_test.go:14 fnlog_test.Example(): i: print infos
	// 2015/06/09 14:32:59 unit_test.go:15 fnlog_test.Example(): w: print warnnings
	// 2015/06/09 14:32:59 unit_test.go:16 fnlog_test.Example(): e: print errors
	// 2015/06/09 14:32:59 unit_test.go:17 fnlog_test.Example(): print messages without tag
}

func init() {
	//noTagLog = fnlog.New("")
}
