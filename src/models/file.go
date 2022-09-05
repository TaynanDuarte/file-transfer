package models

import (
	"bufio"
	"os"
	"strings"
)

type File struct {
	reader              *bufio.Reader
	writer              *bufio.Writer
	originFileReference *os.File
}

func (f *File) FileName() string {

	if strings.Contains(f.originFileReference.Name(), "/") {
		lastSlashIndex := strings.LastIndex(f.originFileReference.Name(), "/")
		return f.originFileReference.Name()[lastSlashIndex+1:]
	}

	return f.originFileReference.Name()
}
