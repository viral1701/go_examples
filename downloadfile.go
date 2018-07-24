package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadfile(uri, filepath string) error {

	outfile, err := os.Create("filepath")
	if err != nil {
		return err
	}

	defer outfile.Close()

	//Downloading File
	fmt.Println("Downloading File")
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
