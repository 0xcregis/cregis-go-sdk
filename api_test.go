package cregis_sdk_go

import (
	"testing"
)

func TestAPI_AddressLegal(t *testing.T) {
	c := NewClient("http://a0c1369e-12ec-467f-9989-7aba384a25e3.apple806.cc:81", "a4b0e563414a4e4dbeb407c89ce2f127", 1388205706190848)
	r, err := c.AddressLegal("195", "TXsmKpEuW7qWnXzJLGP9eDLvWPR2GRn1FS")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}

func TestAPI_AddressInner(t *testing.T) {
	c := NewClient("http://a0c1369e-12ec-467f-9989-7aba384a25e3.apple806.cc:81", "a4b0e563414a4e4dbeb407c89ce2f127", 1388205706190848)
	r, err := c.AddressInner("195", "TXsmKpEuW7qWnXzJLGP9eDLvWPR2GRn1FS")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}

func TestAPI_ListCoins(t *testing.T) {
	c := NewClient("http://a0c1369e-12ec-467f-9989-7aba384a25e3.apple806.cc:81", "a4b0e563414a4e4dbeb407c89ce2f127", 1388205706190848)
	r, err := c.ListCoins()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}

func TestAPI_PayoutQuery(t *testing.T) {
	c := NewClient("http://a0c1369e-12ec-467f-9989-7aba384a25e3.apple806.cc:81", "a4b0e563414a4e4dbeb407c89ce2f127", 1388205706190848)
	r, err := c.PayoutQuery(1390338015477760)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}

func TestAPI_Payout(t *testing.T) {
	c := NewClient("http://a0c1369e-12ec-467f-9989-7aba384a25e3.apple806.cc:81", "a4b0e563414a4e4dbeb407c89ce2f127", 1388205706190848)
	r, err := c.Payout("TXsmKpEuW7qWnXzJLGP9eDLvWPR2GRn1FS", "195@195", "1.1", "c9231e604da54469a735af3f449c880b", "http://xxx.com/payout/callback", "payout")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}
