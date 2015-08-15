package main

import (
	"fmt"
	fs "./filestructure"
	"os"
)

var root *fs.FileStruct

func main() {
	var err error

	root, err = fs.NewDirectoryStruct("root")
	check(err)

	insert("hello world")
	insert("dir1/dir2/")

	// FIXME: this overwrites previous dir1 and dir2 is gone
	insert("dir1")

	root.Print()
}

func FilePath(path string) *fs.FilePath {
	fp, err := fs.NewFilePath(path)
	check(err)
	return fp
}

func insert(path string) {
	_, err := root.InsertFilePath(FilePath(path))
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
