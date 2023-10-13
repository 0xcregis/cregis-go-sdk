package cregis_sdk_go

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/0xcregis/cregis-sdk-go/util"
	"github.com/tidwall/gjson"
)

const (
	URL_PREFIX     = "/api/v1"
	PAYOUT         = URL_PREFIX + "/payout"
	PAYOUT_QUERY   = URL_PREFIX + "/payout/query"
	ADDRESS_LEGAL  = URL_PREFIX + "/address/legal"
	ADDRESS_CREATE = URL_PREFIX + "/address/create"
	ADDRESS_INNER  = URL_PREFIX + "/address/inner"
	PROJECT_COINS  = URL_PREFIX + "/coins"
)

type API struct {
	/**
	 * address
	 */
	url string
	/**
	 * key
	 */
	apiKey string
	/**
	 * projectId
	 */
	pid int64
}

// NewClient user can create a client by the method
/**
  url string host of cregis server
  apiKey string special key to different users
  pid int64 projectId
*/
func NewClient(url, apiKey string, pid int64) *API {
	return &API{
		url: url, apiKey: apiKey, pid: pid,
	}
}

func generateRandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charsetLength := len(charset)
	randomString := make([]byte, length)
	for i := 0; i < length; i++ {
		randomString[i] = charset[rand.Intn(charsetLength)]
	}
	return string(randomString)
}

func (a *API) rebuildParams(params map[string]any) {
	// 生成6位随机字符串
	randomStr := generateRandomString(6)
	params["timestamp"] = time.Now().UnixMilli()
	params["nonce"] = randomStr
}

// AddressLegal Check the legality of the address format
/**
  chainId string code of blockchain
  address string an address of blockchain
*/
func (a *API) AddressLegal(chainId string, address string) (*Result[AddressLegal], error) {
	mp := make(map[string]any, 5)
	mp["pid"] = a.pid
	mp["address"] = address
	mp["chain_id"] = chainId
	a.rebuildParams(mp)
	sign, err := util.DoSign(mp, a.apiKey)
	if err != nil {
		return nil, err
	}
	mp["sign"] = sign
	bs, _ := json.Marshal(mp)
	url := fmt.Sprintf("%v%v", a.url, ADDRESS_LEGAL)
	resp, err := a.send(url, string(bs))
	if err != nil {
		return nil, err
	}
	r := Result[AddressLegal]{}
	err = json.Unmarshal([]byte(resp), &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

// AddressInner whether the address is an internal address
/**
  chainId string code of blockchain
  address string an address of blockchain
*/
func (a *API) AddressInner(chainId, address string) (*Result[AddressInner], error) {
	mp := make(map[string]any, 5)
	mp["pid"] = a.pid
	mp["address"] = address
	mp["chain_id"] = chainId
	a.rebuildParams(mp)
	sign, err := util.DoSign(mp, a.apiKey)
	if err != nil {
		return nil, err
	}
	mp["sign"] = sign
	bs, _ := json.Marshal(mp)
	url := fmt.Sprintf("%v%v", a.url, ADDRESS_INNER)
	resp, err := a.send(url, string(bs))
	if err != nil {
		return nil, err
	}
	r := Result[AddressInner]{}
	err = json.Unmarshal([]byte(resp), &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

// ListCoins currencies supported by the project
func (a *API) ListCoins() (*Result[ProjectCoins], error) {
	mp := make(map[string]any, 5)
	mp["pid"] = a.pid
	a.rebuildParams(mp)
	sign, err := util.DoSign(mp, a.apiKey)
	if err != nil {
		return nil, err
	}
	mp["sign"] = sign
	bs, _ := json.Marshal(mp)
	url := fmt.Sprintf("%v%v", a.url, PROJECT_COINS)
	resp, err := a.send(url, string(bs))
	if err != nil {
		return nil, err
	}
	r := Result[ProjectCoins]{}
	err = json.Unmarshal([]byte(resp), &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

// Payout send tx of withdraw
/**
  address string an address of tx toAddress
  currency string coin flag,like this format 195@195
  amount string value of tx
  thirdPartyId string the business number of user,it is used to callbackUrl
  callbackUrl string set url to receive server event
  remark string to use other input data
*/
func (a *API) Payout(address, currency, amount, thirdPartyId, callbackUrl, remark string) (*Result[Payout], error) {
	mp := make(map[string]any, 5)
	mp["pid"] = a.pid
	mp["currency"] = currency
	mp["address"] = address
	mp["amount"] = amount
	mp["remark"] = remark
	mp["third_party_id"] = thirdPartyId
	mp["callback_url"] = callbackUrl
	a.rebuildParams(mp)
	sign, err := util.DoSign(mp, a.apiKey)
	if err != nil {
		return nil, err
	}
	mp["sign"] = sign
	bs, _ := json.Marshal(mp)
	url := fmt.Sprintf("%v%v", a.url, PAYOUT)
	resp, err := a.send(url, string(bs))
	if err != nil {
		return nil, err
	}
	r := Result[Payout]{}
	err = json.Unmarshal([]byte(resp), &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

// PayoutQuery query order of withdrawal tx
/**
  cid int64 orderId
*/
func (a *API) PayoutQuery(cid int64) (*Result[PayoutQuery], error) {
	mp := make(map[string]any, 5)
	mp["pid"] = a.pid
	mp["cid"] = cid
	a.rebuildParams(mp)
	sign, err := util.DoSign(mp, a.apiKey)
	if err != nil {
		return nil, err
	}
	mp["sign"] = sign
	bs, _ := json.Marshal(mp)
	url := fmt.Sprintf("%v%v", a.url, PAYOUT_QUERY)
	resp, err := a.send(url, string(bs))
	if err != nil {
		return nil, err
	}
	r := Result[PayoutQuery]{}
	err = json.Unmarshal([]byte(resp), &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

// AddressCreate create an address to income
/**
  chainId string code of blockchain
  callbackUrl string set url to receive server event
  alias  string nickname
*/
func (a *API) AddressCreate(chainId, callbackUrl, alias string) (*Result[ProjectAddress], error) {
	mp := make(map[string]any, 5)
	mp["pid"] = a.pid
	mp["chain_id"] = chainId
	mp["alias"] = alias
	mp["callback_url"] = callbackUrl
	a.rebuildParams(mp)
	sign, err := util.DoSign(mp, a.apiKey)
	if err != nil {
		return nil, err
	}
	mp["sign"] = sign
	bs, _ := json.Marshal(mp)
	url := fmt.Sprintf("%v%v", a.url, ADDRESS_CREATE)
	resp, err := a.send(url, string(bs))
	if err != nil {
		return nil, err
	}
	r := Result[ProjectAddress]{}
	err = json.Unmarshal([]byte(resp), &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (a *API) send(url string, payload string) (resp string, err error) {
	body := strings.NewReader(payload)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	r, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if !gjson.ParseBytes(r).Get("code").Exists() {
		return "", fmt.Errorf("%v", string(r))
	}
	return string(r), nil
}
