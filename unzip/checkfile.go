package main

import (
	"fmt"
	"os"
)

func checkfile(src string) error {
	if _, err := os.Stat(src); os.IsNotExist(err) {
		fmt.Println("File Does Not Exist")
	} else {
		deletefile(src)
	}

	return nil
}

func deletefile(src string) error {
	var err = os.Remove(src)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("deleting file")
	return nil
}
