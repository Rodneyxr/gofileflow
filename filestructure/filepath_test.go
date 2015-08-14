package filestructure

import (
	"testing"
)

var invalid_filepaths = [...]string { `/root`, `\root`, `//\/\dir1/file`, `\.file1` }
var valid_filepaths = [...]string { `dir1/`, `file1`, `dir2/file2`, `dir1/file1/`, `dir3/file3//`, `dir4/file//\`,
	`dir1\file2\`, `dir1234567890999/dir1337/file256`, `dir///\/\/\\file1///`, `./file1`, `./dir1/file1`, `dir1/.hidden`,
	`./.hiddendir/.hiddenfile`, `dir1/dir2 with a space/file name`, `hello world`, ` file1`, `  ./dir1`}

func TestNewFilePath(t *testing.T) {
	// create new filepaths from valid filepaths
	for _, path := range valid_filepaths {
		fp, err := NewFilePath(path)
		if err != nil {
			t.Errorf("fail: '%s' should be valid but returned %s", path, fp)
		} else {
			t.Logf("success: '%s' from '%s'", fp, path)
		}
	}

	// create new filepaths from invalid filepaths
	for _, path := range invalid_filepaths {
		fp, err := NewFilePath(path)
		if err != nil {
			t.Logf("success: '%s' was invalid", path)
		} else {
			t.Errorf("fail: '%s' from '%s': should be invalid", fp, path)
		}
	}
}

