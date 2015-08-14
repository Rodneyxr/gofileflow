package filestructure

import (
	"testing"
)

var invalid_filenames = [...]string { "", ".", ".." }
var valid_filenames = [...]string { "root", "/", "1337", " root" }

func TestNewFileStruct(t *testing.T) {
	// create FileStructs from valid filenames
	for _, filename := range valid_filenames {
		fs, err := NewFileStruct(filename)
		if err != nil {
			t.Errorf("fail: '%s' should be valid but returned %s", filename, fs)
		} else {
			t.Logf("success: '%s' from '%s'", fs, filename)
		}
	}

	// create FileStructs from invalid filenames
	for _, filename := range invalid_filenames {
		fs, err := NewFileStruct(filename)
		if err != nil {
			t.Logf("success: '%s' was invalid", filename)
		} else {
			t.Errorf("fail: '%s' from '%s': should be invalid", fs, filename)
		}
	}

}
