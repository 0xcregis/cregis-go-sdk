package main

import (
	"fmt"

	sdk "github.com/0xcregis/cregis-sdk-go"
)

func main() {
	c := sdk.NewClient("http://a0c1369e-12ec-467f-9989-7aba384a25e3.apple806.cc:81", "a4b0e563414a4e4dbeb407c89ce2f127", 1388205706190848)
	r, err := c.AddressLegal("195", "TXsmKpEuW7qWnXzJLGP9eDLvWPR2GRn1FS")
	if err != nil {
		fmt.Printf("err:%v", err.Error())
	} else {
		fmt.Printf("%+v", r)
	}
}
