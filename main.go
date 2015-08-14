package main

import (
	"fmt"
	"./filestructure"
	"os"
)

func main() {
	fs, err := filestructure.NewFileStruct("root")
	check(err)
	fs.Print()

	fp, err := filestructure.NewDirectory("./dir1/dir2/dir2/")
	check(err)
	fmt.Println(fp)

	pathtofp := fp.PathToFile()
	fmt.Println(pathtofp)

	tok := fp.Tokens()
	fmt.Println(tok)

	filename := fp.FileName()
	fmt.Println(filename)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
