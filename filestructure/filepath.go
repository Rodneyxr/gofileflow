package filestructure

import (
	"path/filepath"
	"strings"
	"fmt"
)

type FilePath struct {
	filepath string
	isdir bool
}

// Creates a new FilePath. If the filepath ends with a '/' it will be marked as a directory.
func NewFilePath(path string) (*FilePath, error) {
	fp := new(FilePath)
	path = strings.Replace(path, `\`, `/`, -1)
	fp.filepath = strings.TrimRight(path, `/\`)
	if fp.filepath != path {
		fp.isdir = true
		fmt.Println("isdir = true")
	}
	var err error
	if fp.filepath, err = filepath.Rel("./", path); err != nil { return nil, err }
	return fp, nil
}

func NewDirectory(path string) (*FilePath, error) {
	return NewFilePath(path + "/")
}

// TODO: PathToFile() FilePath
// TODO: Tokenize() string[]
// TODO: FileName() string

func (fp FilePath) FilePath() string {
	return fp.filepath
}

func (fp FilePath) String() string {
	return fp.filepath
}

