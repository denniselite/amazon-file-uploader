package models

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

type File struct {
	Size   int64
	Type   string
	Info   os.FileInfo
	Bytes  []byte
	Basic  *os.File
	Buffer *bufio.Reader
	Name   string
}

func (f *File) Init() file {
	return make(file)
}

func (f *File) Open() (error, uint8, string) {
	var err error
	f.Basic, err = os.Open(fileName)
	if err != nil {
		return err, 4, f.Name
	} else {
		return nil, 0, ""
	}
}

func (f *File) Read() (error, uint8, string) {
	var err error
	f.Info, err = f.Basic.Stat()
	f.Size = f.Info.Size()
	f.Bytes = make([]byte, f.Size)
	f.Buffer = bufio.NewReader(f.Basic)
	_, err = f.Buffer.Read(f.Bytes)
	if err != nil {
		return err, 5, f.Name
	}
	f.Type = http.DetectContentType(f.Bytes)
	return nil, 0, ""
}
