package main

/*
#cgo CFLAGS: -I ./
#cgo LDFLAGS: -L. -lNumPrinter
#include <numPrinter.h>
*/
import "C"

func main() {
	x := 0
	for x < 2 {
	C.print(5);
x++
}
}
