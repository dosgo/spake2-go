package main

import (
	"fmt"
	"spake2-go/ed25519"
)

func main() {
	//var output [300]byte
	//var bytes [64]uint8
	xx := []byte
	arg1 := ed25519.Fiat25519FieldElement{15540839, 8128613, 35539908, 18003200, 66970125, 13484806, 53323836, 19064896, 18909259, 33172296}

	out := ed25519.Fiat25519ToBytes(xx, &arg1)
	fmt.Printf("out:%v\r\n", out)
	return
}
