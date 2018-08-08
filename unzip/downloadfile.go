package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
)

func downloadfile(uri, filepath string) error {

	outfile, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer outfile.Close()

	//Downloading File
	fmt.Println("Downloading File From......" + uri)
	response, err := http.Get(uri)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	_, err = io.Copy(outfile, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func downloadversion(version, path string) error {
	if runtime.GOOS == "windows" {

		err := downloadfile("https://releases.hashicorp.com"+"/terraform/"+version+"/terraform_"+version+"_windows_amd64.zip", path)
		if err != nil {
			panic(err)
		}

	} else if runtime.GOOS == "darwin" {
		err := downloadfile("https://releases.hashicorp.com"+"/terraform/"+version+"/terraform_"+version+"_darwin_amd64.zip", path)
		if err != nil {
			panic(err)
		}
	} else if runtime.GOOS == "linux" {
		err := downloadfile("https://releases.hashicorp.com"+"/terraform/"+version+"/terraform_"+version+"_linux_amd64.zip", path)
		if err != nil {
			panic(err)
		}

	}

	return nil

}
