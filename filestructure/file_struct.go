package filestructure

import (
	"errors"
)

type FileStruct struct {
	Name string
	files map[string]FileStruct
}

// returns a new FileStruct type
func NewFileStruct(name string) (*FileStruct, error) {
	if name == "" || name == "." || name == ".." {
		return nil, errors.New("FileStruct: Invalid file name.")
	}

	fs := new(FileStruct)
	fs.Name = name
	fs.files = make(map[string]FileStruct)
	fs.files["."] = *fs
	fs.files[".."] = *fs
	return fs, nil
}

func (fs FileStruct) IsNil() bool {
	return fs.Name == ""
}

func (fs FileStruct) String() string {
	return fs.Name
}
