package main

import (
	log "github.com/golang/glog"
	"flag"
	"os"
	"fmt"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARN|FATAL] -log_dir=[string]\n", )
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	flag.Usage = usage
	// NOTE: This next line is key you have to call flag.Parse() for the command line 
	// options or "flags" that are defined in the glog module to be picked up.
	flag.Parse()
}


func main() {
	
	log.Info("Info")
	log.Warning("Warning")
	log.Error("Error")
//	log.Fatal("Fatal")
}

