package rofio

import (
	"os"
)

type FILE_OPEN_FUNC func (string) (*os.File, error)

func open_file (f_name string, f_func FILE_OPEN_FUNC) (*os.File, func()) {
	fi, err := f_func (f_name)
	if (err != nil) {
		panic (err)
	}
	return fi, func() { if err := fi.Close(); err != nil { panic (err) } }
}

func Open_file (file_name string) (*os.File, func()) {
	return open_file (file_name, os.Open)
}

func Create_file (file_name string) (*os.File, func()) {
	return open_file (file_name, os.Create)
}
