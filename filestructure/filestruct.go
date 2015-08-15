package filestructure

import (
	"errors"
	"strings"
	"fmt"
)

type FileStruct struct {
	Name   string
	parent *FileStruct
	files  map[string]*FileStruct
	isdir  bool
}

// returns a new FileStruct type
func NewFileStruct(name string) (*FileStruct, error) {
	name = strings.TrimSpace(name)
	if name == "" || name == "." || name == ".." {
		return nil, errors.New("FileStruct: Invalid file name.")
	}

	fs := new(FileStruct)
	fs.Name = name
	fs.parent = fs
	fs.files = nil // files can't hold files
	fs.isdir = false

	return fs, nil
}

// returns a new FileStruct that is a directory
func NewDirectoryStruct(name string) (*FileStruct, error) {
	var fs *FileStruct
	var err error
	if fs, err = NewFileStruct(name); err != nil { return fs, err }
	fs.isdir = true
	fs.files = map[string]*FileStruct{
		"." : fs,
		".." : fs.parent,
	}
	return fs, nil
}

// insert a FilePath into the FileStruct creating any directories that do not exist.
func (fs *FileStruct) InsertFilePath(fp *FilePath) (*FileStruct, error) {
	if fs == nil {
		return nil, errors.New("FilePath: cannot be nil")
	}
	// first check if the FileStruct is a directory
	if !fs.isdir {
		return nil, errors.New(fmt.Sprintf("FilePath: cannot create directory '%s': not a directory", fp))
	}

	// create the full directory path or path to the file
	newfs, err := fs.mkdir(fp)
	if err != nil {
		return nil, err
	}

	// if the filepath was a directory, mkdir only make the path to the file
	// so the file needs to be inserted into the directory returned by mkdir()
	if !fp.isdir {
		newfs, err = newfs.insert(fp.FileName(), false)
		if err != nil {
			return nil, err
		}
	}
	return newfs, nil
}

// creates all necessary directories to the specified filepath
func (fs *FileStruct) mkdir(fp *FilePath) (*FileStruct, error) {
	var err error
	safe := false // will be flagged true if it is safe to create the rest of the path without checking if it exists
	cp := fs // save current pointer
	for _, filename := range fp.PathToFile().Tokens() {
		// peek ahead to check if the next level exists
		if !safe {
			if peek := cp.files[filename]; peek != nil {
				// if the next level exists but is not a directory return an error
				if !peek.isdir {
					return nil, errors.New(fmt.Sprintf("FileStruct: cannot create directory '%s': not a directory", fp))
				} else {
					cp = peek
				}
			} else {
				cp, err = cp.insert(filename, true)
				if err != nil {
					return nil, err
				}
			}
		} else {
			// else if it does not exist we create it and stop checking if next levels exist
			cp, err = cp.insert(filename, true)
			if err != nil {
				return nil, err
			}
			safe = true
		}
	}
	return cp, nil
}

func (fs *FileStruct) insert(filename string, isdir bool) (*FileStruct, error) {
	var tmp *FileStruct
	var err error

	if isdir {
		tmp, err = NewDirectoryStruct(filename)
	} else {
		tmp, err = NewFileStruct(filename)
	}
	if err != nil { return nil, err }

	tmp.setParent(fs)
	return tmp, nil
}

func (child *FileStruct) setParent(parent *FileStruct) {
	child.parent = parent
	if child.isdir {
		child.files[".."] = parent
	}
	parent.files[child.Name] = child
}

func (fs *FileStruct) FileExists(fp *FilePath) *FileStruct {
	cp := fs // save current pointer
	for _, filename := range fp.Tokens() {
		if cp := cp.files[filename]; cp == nil {
			return nil
		}
	}
	return cp
}

// print the FileStruct
func (fs *FileStruct) Print() {
	fs.printFileStruct(0)
}

// recursively traverse the tree and print each node representing a directory structure
func (fs *FileStruct) printFileStruct(level int) {
	// fmt.Printf("Level %d:\t%s%s\n", level, strings.Repeat("\t", level), fs.DisplayName())
	fmt.Printf("%s%s\n", strings.Repeat("\t", level), fs.DisplayName())

	// iterate over all files in the file struct
	for _, file := range fs.files {
		// ignore the directory if itself or parent
		if file == fs || file == fs.parent { continue }
		file.printFileStruct(level + 1)
	}
}

// return true if the FileStruct is nil
func (fs *FileStruct) IsNil() bool {
	return fs.Name == ""
}

// returns the name as a string with an ending slash if isdir == true
func (fs *FileStruct) DisplayName() string {
	if fs.isdir {
		return fs.Name + "/"
	}
	return fs.Name
}

func (fs FileStruct) String() string {
	return fs.Name
}
