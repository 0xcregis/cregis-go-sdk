package cregis_sdk_go

type Result[T any] struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

type AddressLegal struct {
	/**
	 * 地址是否合法
	 */
	Result bool `json:"result"`
}

type AddressInner struct {
	/**
	 * 地址是否合法
	 */
	Result bool `json:"result"`
}

type ProjectCoin struct {
	CoinName string `json:"coin_name"`
	ChainId  string `json:"chain_id"`
	TokenId  string `json:"token_id"`
}

type ProjectCoins struct {
	PayoutCoins  []*ProjectCoin `json:"payout_coins"`
	AddressCoins []*ProjectCoin `json:"address_coins"`
	OrderCoins   []*ProjectCoin `json:"order_coins"`
}

type Payout struct {
	/**
	 * 转出订单编号
	 */
	Cid int64 `json:"cid"`
}

type PayoutQuery struct {
	/**
	 * 项目编号
	 */
	Pid int64 `json:"pid"`
	/**
	 * 地址
	 */
	Address string `json:"address"`
	/**
	 * 链编号
	 */
	ChainId string `json:"chain_id"`
	/**
	 * 代币编号
	 */
	TokenId string `json:"token_id"`
	/**
	 * 币种标识
	 */
	Currency string `json:"currency"`
	/**
	 * 金额
	 */
	Amount string `json:"amount"`
	/**
	 * 调用方业务编号
	 */
	ThirdPartyId string `json:"third_party_id"`
	/**
	 * 备注
	 */
	Remark string `json:"remark"`
	/**
	 * 状态
	 */
	Status int `json:"status"`
	/**
	 * 交易哈希
	 */
	TxId string `json:"txid"`
	/**
	 * 区块高度
	 */
	BlockHeight string `json:"block_height"`
	/**
	 * 区块时间
	 */
	BlockTime string `json:"block_time"`
}

type ProjectAddress struct {
	Address string `json:"address"`
}
