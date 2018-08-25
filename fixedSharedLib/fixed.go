package main

import (
	"fmt"
)

type message string

func (m message) DoSomething(name string) {
	fmt.Println("Sent " + name + " to X - fixed version lib")
}

// Machine - Exporting Symbol
var Machine message
