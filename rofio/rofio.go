package rofio

import (
	"os"
)

type FILE_OPEN_FUNC func (string) (*os.File, error)

// This internal file open function opens a file using the specified
// filename and specific open function from the OS package.
// It also returns an anonymous function for closing the file which is
// meant to be deferred at the call site.
func open_file (f_name string, f_func FILE_OPEN_FUNC) (*os.File, func()) {
	fi, err := f_func (f_name)
	if (err != nil) {
		panic (err)
	}
	return fi, func() { if err := fi.Close(); err != nil { panic (err) } }
}

// Open an existing file.
func Open_file (file_name string) (*os.File, func()) {
	return open_file (file_name, os.Open)
}

// Create a new file.
func Create_file (file_name string) (*os.File, func()) {
	return open_file (file_name, os.Create)
}
