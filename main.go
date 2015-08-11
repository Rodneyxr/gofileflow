package main

import (
	"fmt"
	"github.com/rodneyxr/gofileflow/filestructure"
)

func main() {
	fs := filestructure.FileStruct{"testing"}
	fmt.Println(fs.String())
}
