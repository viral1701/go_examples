package main

import (
	"flag"
	"log"
)

func main() {

	version := flag.String("version", "", "versionofterraform")
	//zipfile := flag.String("zipfile", "", "zipfilepath")
	//extract := flag.String("extract", "", "extractdestination")
	path := flag.String("path", "", "downloadpath")
	flag.Parse()

	err := downloadversion(*version, *path)
	if err != nil {
		log.Fatal(err)
	}

}
