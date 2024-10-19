# spake2-go



func main() {
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
	if status < 1 {
		log.Printf("Unable to process their public key")
		return
	}

	log.Printf("ALICE(%d) <== %v\n", aKeyLen, aKey)
	log.Printf("BOB(%d)  <== %v\n", bKeyLen, bKey)
	return
}
