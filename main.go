package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/Masterminds/log-go"
)

const out = `Hello Rancher Desktop`

var debug = flag.Bool("debug", false, "display debug output")

func handler(w http.ResponseWriter, r *http.Request) {
	log.Info("Printing message")
	log.Debugf("Request %v", r)
	fmt.Fprint(w, out)
}

func main() {

	flag.Parse()
	logger := log.NewStandard()
	if *debug {
		logger.Level = log.DebugLevel
	}
	log.Current = logger

	log.Info("Starting Hello World")

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
