package filestructure

import (

)

type FileStruct struct {
	Name string
	Files map[string]FileStruct
}

func (fs FileStruct) String() string {
	return fs.Name
}
