
# fnlog

#### About

fnlog is a helper for golang `log` package to add tag, file name, function name, line number while printing log message.  
It make it easier to debug golang applications.

    import (
        "github.com/northbright/fnlog"
        "log"
    )

    func main() {
        iLog := fnlog.New("i")
        wLog := fnlog.New("w")
        eLog := fnlog.New("e")
        noTagLog := fnlogNew("")

        iLog.Printf("print infos")
        wLog.Printf("print warnnings")
        eLog.Printf("print errors")
        newLog.Printf("print messages without tag")    
    }

    // Output:
    // ===============================================================================
    2015/04/24 10:43:25 unit_test.go:16 fnlog.TestLogger(): i: print infos
    2015/04/24 10:43:25 unit_test.go:17 fnlog.TestLogger(): w: print warnnings
    2015/04/24 10:43:25 unit_test.go:18 fnlog.TestLogger(): e: print errors
    2015/04/24 10:43:25 unit_test.go:19 fnlog.TestLogger(): print messages without tag

#### Thanks

Many thanks to <https://github.com/wyc/funclog>  
I only made a little changes for my own project:  

* Keep package name while printing function name
* Add TAG check to fill the prefix of log

####  License

MIT License

