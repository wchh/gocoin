package main

import (
	"encoding/hex"
	"github.com/wchh/gocoin/lib/others/cgo/openssl"
	"time"
)

var CNT int = 15 * 250

func main() {
	key, _ := hex.DecodeString("040eaebcd1df2df853d66ce0e1b0fda07f67d1cabefde98514aad795b86a6ea66dbeb26b67d7a00e2447baeccc8a4cef7cd3cad67376ac1c5785aeebb4f6441c16")
	sig, _ := hex.DecodeString("3045022100fe00e013c244062847045ae7eb73b03fca583e9aa5dbd030a8fd1c6dfcf11b1002207d0d04fed8fa1e93007468d5a9e134b0a7023b6d31db4e50942d43a250f4d07c01")
	msg, _ := hex.DecodeString("3382219555ddbb5b00e0090f469e590ba1eae03c7f28ab937de330aa60294ed6")
	sta := time.Now()
	for i := 0; i < CNT; i++ {
		if openssl.EC_Verify(key, sig, msg) != 1 {
			println("Verify error")
			return
		}
	}
	sto := time.Now()
	println((sto.UnixNano()-sta.UnixNano())/int64(CNT*1000), "us per ECDSA_Verify")
}
