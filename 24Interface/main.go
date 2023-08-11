// interface is a contract
// golang, no need to enforce interface
//in golang if a type satisfies an interface

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	c := new(Console)
	fw := new(FW)
	fw.FileName = "log.txt"

	// n, err := c.Write([]byte("Hello"))
	// if err != nil {
	// 	fmt.Println(err)
	// }

	fmt.Fprintln(c, "Hello World")
	fmt.Fprintln(fw, "Hello World")
}

type Console struct{}

func (c *Console) Write(bytes []byte) (n int, err error) {
	return fmt.Println(string(bytes))
	//return 0, nil
}

type FW struct {
	FileName string
}

func (f *FW) Write(bytes []byte) (n int, err error) {
	if f == nil {
		return 0, errors.New("nil file object")
	}
	f1, err := os.Open(f.FileName)
	if err != nil {
		return 0, err
	}
	return f1.Write(bytes)
}
