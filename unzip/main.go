package main

import (
	"flag"
)

func main() {

	version := flag.String("version", "", "versionofterraform")
	extract := flag.String("extract", "", "where it extracts")
	flag.Parse()

	downloadversion(*version, *extract)

}
