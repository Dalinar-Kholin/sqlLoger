package main

import (
	"fmt"
	"os"
)

var (
	fileName = ""
	handler  *os.File
)

func Exec(s ...string) {
	var r string
	for _, v := range s {
		r += v
	}
	handler.Write([]byte(r + "\n"))
	fmt.Printf("%v\n", r)
}

func SetUpLogger(file string) error {
	var err error
	handler, err = os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	return err
}

func essa[T float32 | int](xd T) {

}
