package main

import (
	"flag"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()

	glog.Info("Starting up ssh-trap")

	glog.Info("Closing ssh-trap")
	glog.Flush()
}