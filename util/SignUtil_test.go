package util

import (
	"testing"
)

func TestDoSign(t *testing.T) {

	//	str := `
	//{
	//  "pid": 1382528827416576,
	//  "currency": "195@195",
	//  "address": "TXsmKpEuW7qWnXzJLGP9eDLvWPR2GRn1FS",
	//  "amount": "1.1",
	//  "remark": "payout",
	//  "third_party_id": "c9231e604da54469a735af3f449c880f",
	//  "callback_url": "http://192.168.2.29:9099/callback",
	//  "nonce": "hwlkk6",
	//  "timestamp": 1688004243314
	//}
	//`

	mp := map[string]any{"pid": 1382528827416576, "address": "TXsmKpEuW7qWnXzJLGP9eDLvWPR2GRn1FS"}
	s, err := DoSign(mp, "f502a9ac9ca54327986f29c03b271491")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(s)
	}

}
