package main

import (
	"fmt"
	"io"
)

const englishGreetingPrefix = "Hello, "

func Greet(w io.Writer, name string) {
	fmt.Fprint(w, englishGreetingPrefix+name)
}
