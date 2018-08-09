package main

import (
	"flag"
)

func main() {

	version := flag.String("version", "", "versionofterraform")
	//zipfile := flag.String("zipfile", "", "zipfilepath")
	extract := flag.String("extract", "", "extractdestination")
	path := flag.String("path", "", "downloadpath")
	flag.Parse()

	downloadversion(*version, *path)
	unzip(*path, *extract)

}
