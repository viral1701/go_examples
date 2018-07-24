package main

import (
	"fmt"
	"os"
)

func deletefile(src string) error {
	var err = os.Remove(src)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("deleting file")
	return nil
}
