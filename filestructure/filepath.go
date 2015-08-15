package filestructure

import (
	"path/filepath"
	"strings"
	"regexp"
)

type FilePath struct {
	filepath string
	isdir bool
}

// Creates a new FilePath. If the filepath ends with a '/' it will be marked as a directory.
func NewFilePath(path string) (*FilePath, error) {
	fp := new(FilePath)
	path = strings.TrimSpace(path)
	path = strings.Replace(path, `\`, `/`, -1)
	fp.filepath = strings.TrimRight(path, `/`)
	if fp.filepath != path {
		fp.isdir = true
	}
	var err error
	if fp.filepath, err = filepath.Rel("./", path); err != nil { return nil, err }
	return fp, nil
}

// returns a FilePath that is a directory
func NewDirectoryPath(path string) (*FilePath, error) {
	return NewFilePath(path + "/")
}

// returns the a filepath to the file (excluding the last element)
// if the filepath is points to a directory then that directory (self) will be returned
// if an error occurs (should never happen) nil will be returned
func (fp *FilePath) PathToFile() *FilePath {
	if fp.isdir { return fp }
	path, err := NewFilePath(filepath.Dir(fp.filepath))
	if err != nil { return nil }
	return path
}

// splits the file path into separate strings representing each level
func (fp *FilePath) Tokens() []string {
	return regexp.MustCompile(`/|\\`).Split(fp.filepath, -1)
}

// returns the last file in the path
func (fp *FilePath) FileName() string {
	return filepath.Base(fp.filepath)
}

// returns true if the filepath points to a directory
func (fp *FilePath) IsDir() bool {
	return fp.isdir
}

// return a string representing the full path
func (fp *FilePath) FilePath() string {
	return fp.filepath
}

func (fp FilePath) String() string {
	return fp.filepath
}

