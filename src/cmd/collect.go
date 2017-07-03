package main

import (
	"runtime"
	"time"
	"math/rand"
	"cmd/app"
	_"model/collect"
	_"dao/sql_reg"
	"log"
	"net/http"
	"control"
	"common"
)

func main() {
	if common.Debug0 == true {
		common.DebugPrint("The cpu core is", runtime.NumCPU(), ",The app would use all of cores")
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UTC().UnixNano())
	go app.Run()
	log.Fatal(http.ListenAndServe(":8080", func() http.Handler {
		s, _ := control.CollectRouters()
		return s
	}()))
}
