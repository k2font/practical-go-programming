package main

import (
	"fmt"

	"golang.org/x/xerrors"
)

func StackTrace() {
	err := xerrors.New("error")
	fmt.Printf("%+v\n", err)
}
