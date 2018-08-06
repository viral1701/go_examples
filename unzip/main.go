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

	downloaderr := downloadversion(*version, *path)
	if downloaderr != nil {
		log.Fatal(downloaderr)
	}

}
