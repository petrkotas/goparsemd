package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
)

var dataPath string

// also can use ioutil.ReadFile(filename)
func readFile(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}

// main function is the entrypoint of an application
func main() {
	log.Printf("Data path %s", dataPath)

	files, err := ioutil.ReadDir(dataPath)
	if err != nil {
		log.Fatalf("Cannot load data, %s", err.Error())
	}

	for _, file := range files {
		log.Printf("File: %s", file.Name())
	}

}

// init function is executed first,
// normal use is initilization of the package
// only one init can be in the package
func init() {
	flag.StringVar(&dataPath, "data", "", "set data source directory")
	flag.Parse()
}
