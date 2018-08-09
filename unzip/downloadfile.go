package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
)

func downloadfile(uri, filepath string) error {

	outfile, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}

	defer outfile.Close()

	//Downloading File
	fmt.Println("Downloading File From......" + uri)
	response, err := http.Get(uri)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer response.Body.Close()

	_, err = io.Copy(outfile, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func downloadversion(version, path string) error {
	if runtime.GOOS == "windows" {

		downloadfile("https://releases.hashicorp.com"+"/terraform/"+version+"/terraform_"+version+"_windows_amd64.zip", path)

	} else if runtime.GOOS == "darwin" {
		downloadfile("https://releases.hashicorp.com"+"/terraform/"+version+"/terraform_"+version+"_darwin_amd64.zip", path)

	} else if runtime.GOOS == "linux" {
		downloadfile("https://releases.hashicorp.com"+"/terraform/"+version+"/terraform_"+version+"_linux_amd64.zip", path)

	}

	return nil

}
