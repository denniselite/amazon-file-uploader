package myErrors

import (
	"errors"
	"fmt"
	"os"
)

type Error struct {
	List    map[uint8]string
	err     error
	key     uint8
	message string
}

func Init() Error {
	var e Error
	e.setList()
	return e
}

func (e *Error) setList() {
	e.List[1] = "Unsupported params : "
	e.List[2] = "Required params is empty : "
	e.List[3] = "File not found : "
	e.List[4] = "Error open file : "
	e.List[5] = "Error read file : "
	e.List[6] = "Can't connect to server : "
	e.List[6] = "Amazon auth failed : "
	e.List[7] = "Connection is lost : "
	e.List[8] = "Add to bucket failed : "
}

func (e *Error) Check(err error, key uint8, msg string) {
	if err != nil {
		fmt.Println(string(e.List[key] + msg))
		fmt.Println(err)
		os.Exit(key)
	}
	if (err == nil) && (key > 0) && (mgs != "") {
		fmt.Println(string(e.List[key] + msg))
		os.Exit(key)
	}
}
