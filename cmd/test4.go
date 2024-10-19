package main

import (
	"fmt"
)

func main() {

	var x22 uint32 = 0
	var x23 uint8 = 0
	var x1 uint32 = 15540858
	var x21 uint32 = 4294967295

	fmt.Printf("x23:%d x1:%d x21:%d ddd:%d\r\n", x23, x1, x21, (x21 & uint32(0x3ffffed)))
	curve25519go.f(&x22, &x23, 0x0, x1, (x21 & uint32(0x3ffffed)))

	fmt.Printf("x22:%v\r\n", x22)
	return
}
