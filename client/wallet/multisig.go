package wallet

import (
	"encoding/json"
	"github.com/wchh/gocoin/client/common"
	"github.com/wchh/gocoin/lib/btc"
	"io/ioutil"
	"os"
)

type MultisigAddr struct {
	MultiAddress               string
	ScriptPubKey               string
	KeysRequired, KeysProvided uint
	RedeemScript               string
	ListOfAddres               []string
}

func IsMultisig(ad *btc.BtcAddr) (yes bool, rec *MultisigAddr) {
	yes = ad.Version == btc.AddrVerScript(common.Testnet)
	if !yes {
		return
	}

	fn := common.CFG.Walletdir +
		string(os.PathSeparator) + "multisig" +
		string(os.PathSeparator) + ad.String() + ".json"

	d, er := ioutil.ReadFile(fn)
	if er != nil {
		//println("fn", fn, er.Error())
		return
	}

	var msa MultisigAddr
	er = json.Unmarshal(d, &msa)
	if er == nil {
		rec = &msa
	} else {
		println(fn, er.Error())
	}

	return
}
