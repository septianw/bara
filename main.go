package main

import (
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/septianw/bara/libs"
	"github.com/septianw/bara/router"
	"github.com/septianw/golok"
)

var Writelog = golok.Writelog

func main() {
	var routers = router.Router()
	runtime.GOMAXPROCS(4)
	*libs.Loglevel = os.Getenv("LOGLEVEL")
	*libs.Logfile = os.Getenv("LOGFILE")

	http.Handle("/", routers)

	Writelog("info", "Server run on localhost:8484")
	log.Fatal(http.ListenAndServe(":8484", nil))
}
