package main

import (
	"fmt"
	"./filestructure"
	"os"
)

func main() {
	fs, err := filestructure.NewFileStruct("root")
	check(err)

	fmt.Println(fs)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
