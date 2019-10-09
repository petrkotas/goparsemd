package main

import (
	"flag"
	"log"
)

var dataPath string

func main() {
	log.Printf("Data path %s", dataPath)
}

// init function is executed first,
// normal use is initilization of the package
// only one init can be in the package
func init() {
	flag.StringVar(&dataPath, "data", "", "set data source directory")
	flag.Parse()
}
