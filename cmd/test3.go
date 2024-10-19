package main

import (
	"fmt"
	"spake2-go/curve25519go"
)

func main() {
	//var output [300]byte
	//var bytes [64]uint8
	out := make([]byte, 32)
	arg1 := [10]uint32{15540839, 8128613, 35539908, 18003200, 66970125, 13484806, 53323836, 19064896, 18909259, 33172296}

	curve25519go.Fiat_25519_to_bytes(&out, &arg1)
	fmt.Printf("out:%v\r\n", out)
	return
}
