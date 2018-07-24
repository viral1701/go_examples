package main

import (
	"flag"
	"log"
)

func main() {
	filetodelete := "c:\\Temp3\\terraform.zip"
	zipfile := flag.String("zipfile", "", "zipfilepath")
	extact := flag.String("extract", "", "extractdestination")
	flag.Parse()

	checkfile(filetodelete)

	err := unzip(*zipfile, *extact)
	if err != nil {
		log.Fatal(err)
	}
}
