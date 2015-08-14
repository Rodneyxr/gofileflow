package filestructure

import (
	"path/filepath"
	"strings"
)

type FilePath struct {
	filepath string
	isdir bool
}

// Creates a new FilePath. If the filepath ends with a '/' it will be marked as a directory.
// FIXME: strip whitespace from the name
func NewFilePath(path string) (*FilePath, error) {
	fp := new(FilePath)
	path = strings.Replace(path, `\`, `/`, -1)
	fp.filepath = strings.TrimRight(path, `/`)
	if fp.filepath != path {
		fp.isdir = true
	}
	var err error
	if fp.filepath, err = filepath.Rel("./", path); err != nil { return nil, err }
	return fp, nil
}

func NewDirectory(path string) (*FilePath, error) {
	return NewFilePath(path + "/")
}

// TODO: PathToFile() FilePath
func (fp *FilePath) PathToFile() *FilePath {
	//if fp.isdir {
		//return fp
	//}
	path, err := NewFilePath(filepath.Dir(fp.filepath))
	if err != nil { return nil }
	return path
}

// TODO: Tokenize() string[]
// TODO: FileName() string

func (fp *FilePath) IsDir() bool {
	return fp.isdir
}

func (fp *FilePath) FilePath() string {
	return fp.filepath
}

func (fp FilePath) String() string {
	return fp.filepath
}

