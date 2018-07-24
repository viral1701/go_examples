package main

import (
	"flag"
	"log"
)

func main() {
	zipfile := flag.String("zipfile", "", "zipfilepath")
	extact := flag.String("extract", "", "extractdestination")
	flag.Parse()

	err := unzip(*zipfile, *extact)
	if err != nil {
		log.Fatal(err)
	}
}
