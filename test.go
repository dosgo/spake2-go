package main

import (
	"log"
	"spake2-go/spake2"
)

// 假设有一个名为 spake2 的包提供了 SPAKE2 协议的实现

const (
	kClientName = "adb pair client"
	kServerName = "adb pair server"
)

var SPAKE2_MAX_MSG_SIZE int = 32
var SPAKE2_MAX_KEY_SIZE = 64

var kPassword = []byte{
	0x35, 0x39, 0x32, 0x37, 0x38, 0x31, 0xe6, 0x3d, 0xd9, 0x59, 0x65, 0x1c,
	0x21, 0x16, 0x00, 0xf3, 0xb6, 0x56, 0x1d, 0x0b, 0x9d, 0x90, 0xaf, 0x09,
	0xd0, 0xa4, 0xa4, 0x53, 0xee, 0x20, 0x59, 0xa4, 0x80, 0xcc, 0x7c, 0x5a,
	0x94, 0xd4, 0xd4, 0x89, 0x33, 0xf9, 0xff, 0xf5, 0xfe, 0x43, 0x31, 0x7d,
	0x52, 0xfa, 0x7b, 0xff, 0x8f, 0x8b, 0xc4, 0xf3, 0x48, 0x8b, 0x80, 0x07,
	0x33, 0x0f, 0xec, 0x7c, 0x7e, 0xdc, 0x91, 0xc2, 0x0e, 0x5d,
}

func hexify(in []byte) string {
	var out []byte
	for _, b := range in {
		out = append(out, "0123456789ABCDEF"[b>>4])
		out = append(out, "0123456789ABCDEF"[b&0x0F])
	}
	return string(out)
}

func main() {
	//var output [300]byte
	//var bytes [64]uint8

	alice, err := spake2.SPAKE2_CTX_new(0, []byte(kClientName), []byte(kServerName))
	bob, err := spake2.SPAKE2_CTX_new(1, []byte(kServerName), []byte(kClientName))
	if alice == nil || bob == nil || err != nil {
		log.Printf("Unable to create a SPAKE2 context.")
		return
	}
	var aMessage = make([]byte, uint32(SPAKE2_MAX_MSG_SIZE))
	var bMessage = make([]byte, uint32(SPAKE2_MAX_MSG_SIZE))
	var aMsgSize uint32 = uint32(SPAKE2_MAX_MSG_SIZE)
	var bMsgSize uint32 = uint32(SPAKE2_MAX_MSG_SIZE)

	status := alice.SPAKE2_generate_msg(aMessage, aMsgSize, kPassword)

	status1 := bob.SPAKE2_generate_msg(bMessage, bMsgSize, kPassword)

	log.Printf("ALICE(%d) ==>:%+v\r\n", aMsgSize, aMessage)
	log.Printf("BOB(%d)   ==>:%+v\r\n", bMsgSize, bMessage)
	if status < 1 || status1 < 1 || aMsgSize == 0 || bMsgSize == 0 {
		log.Printf("Unable to generate the SPAKE2 public key111.")
		return
	}

	var aKeyLen uint32 = uint32(SPAKE2_MAX_KEY_SIZE)
	var aKey = make([]byte, uint32(SPAKE2_MAX_KEY_SIZE))
	status, err = alice.SPAKE2_process_msg(aKey, aKeyLen, bMessage)

	var bKeyLen uint32 = uint32(SPAKE2_MAX_KEY_SIZE)
	var bKey = make([]byte, uint32(SPAKE2_MAX_KEY_SIZE))
	status1, _ = bob.SPAKE2_process_msg(bKey, bKeyLen, aMessage)
	log.Printf("status:%d err:%v\r\n", status, err)
	log.Printf("status1:%d\r\n", status1)
	if status != 2 {
		log.Printf("Unable to process their public key")
		return
	}

	return
}
