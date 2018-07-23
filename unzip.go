package main

import (
	"archive/zip"
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"
)

func unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		filereader, err := file.Open()
		if err != nil {
			return err
		}

		defer filereader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {

			return err
		}

		defer targetFile.Close()

		if _, err := io.Copy(targetFile, filereader); err != nil {
			return err
		}
	}

	return nil

}

func main() {

	zipfile := flag.String("zipfile", "", "zipfilepath")
	extact := flag.String("extract", "", "extractdestination")
	flag.Parse()

	err := unzip(*zipfile, *extact)
	if err != nil {
		log.Fatal(err)
	}

}
