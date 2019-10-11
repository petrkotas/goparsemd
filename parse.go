package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"gopkg.in/yaml.v3"
)

// explain package concept, what is package, how variables works and what is exported

var dataPath string
var outputPath string

// Event information used for conversion
type Event struct {
	Title   string `yaml:"title" json:"title"`
	Date    string `yaml:"event_date" json:"date"`
	Time    string `yaml:"event_time" json:"time"`
	Excerpt string `yaml:"excerpt" json:"excerpt"`
	Content string `json:"content"`
}

// also can use ioutil.ReadFile(filename)
func readFile(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}

func writeFile(file string, data []byte) error {
	log.Printf("Writing to file: %s", file)
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}

	defer func() {
		err := f.Close()
		if err != nil {
			log.Printf("Cannot close the file, %s", err.Error())
		}
	}()

	n, err := f.Write(data)
	if err != nil {
		return err
	}
	if n != len(data) {
		log.Printf("Written only %d bytes of %d", n, len(data))
	}
	return nil
}

// parse the data to the workable data structure
func parseData(data []byte) Event {
	// extract the header information
	re := regexp.MustCompile(`---([\s\S]*?)---`)
	header := re.FindSubmatch(data)[1]

	// header now contains YAML, now what?
	event := Event{}
	err := yaml.Unmarshal(header, &event)
	if err != nil {
		log.Printf("Cannot convert to yaml, %s", err.Error())
	}

	// now parse the content
	re = regexp.MustCompile(`[\s\S]*---\s([\s\S]*)`)
	content := re.FindSubmatch(data)[1]
	event.Content = string(content)

	return event
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
		data, _ := readFile(filepath.Join(dataPath, file.Name()))

		event := parseData(data)

		jsonData, err := json.Marshal(event)
		if err != nil {
			log.Printf("Cannot encode to json, %s", err.Error())
			continue
		}

		writeFile(filepath.Join(outputPath, file.Name()+".json"), jsonData)

	}
}

// init function is executed first,
// normal use is initilization of the package
// only one init can be in the package
func init() {
	flag.StringVar(&dataPath, "data-in", "", "set data source directory")
	flag.StringVar(&outputPath, "data-out", "", "set data source directory")
	flag.Parse()
}
