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

func downloadversion(version, extractpath string) error {
	if runtime.GOOS == "windows" {
		var path string = os.Getenv("Temp") + "terraform.zip"
		downloadfile("https://releases.hashicorp.com"+"/terraform/"+version+"/terraform_"+version+"_windows_amd64.zip", path)
		unzip(path, extractpath)

	} else if runtime.GOOS == "darwin" {

		var path string = os.Getenv("TMPDIR") + "terraform.zip"
		downloadfile("https://releases.hashicorp.com"+"/terraform/"+version+"/terraform_"+version+"_darwin_amd64.zip", path)
		unzip(path, extractpath)

	} else if runtime.GOOS == "linux" {
		var path string = os.Getenv("TMPDIR") + "terraform.zip"
		downloadfile("https://releases.hashicorp.com"+"/terraform/"+version+"/terraform_"+version+"_linux_amd64.zip", path)
		unzip(path, extractpath)

	}

	return nil

}
