package main

import (
	"fmt"
	"github.com/wchh/gocoin/lib"
	"github.com/wchh/gocoin/lib/btc"
	"github.com/wchh/gocoin/lib/others/utils"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Gocoin FetchTx version", lib.Version)

	if len(os.Args) < 2 {
		fmt.Println("Specify transaction id on the command line (MSB).")
		return
	}

	txid := btc.NewUint256FromString(os.Args[1])
	rawtx := utils.GetTxFromWeb(txid)
	if rawtx == nil {
		fmt.Println("Error fetching the transaction")
	} else {
		ioutil.WriteFile(txid.String()+".tx", rawtx, 0666)
	}
}
