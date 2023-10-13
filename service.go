package cregis_sdk_go

type CregisClientInterface interface {
	//AddressLegal
	/**
	 * 检测地址合法性
	 *
	 * @param chainId 链编号
	 * @param address 地址
	 * @return 结果
	 */
	AddressLegal(chainId string, address string) (*Result[AddressLegal], error)

	//AddressInner
	/**
	 * 检测地址在项目中是否存在
	 *
	 * @param chainId 链编号
	 * @param address 地址
	 * @return 结果
	 */
	AddressInner(chainId, address string) (*Result[AddressInner], error)

	//ListCoins
	/**
	 * 获取支持的币种
	 * @return 币种列表
	 */
	ListCoins() (*Result[ProjectCoins], error)

	//Payout
	/**
	 * 发送提币申请
	 *
	 * @param address      提币地址
	 * @param currency 币种信息
	 *                 规则参考 <a href="https://app.apifox.com/link/project/2923699/apis/doc-2804947">https://app.apifox.com/link/project/2923699/apis/doc-2804947</a>
	 * @param amount       金额
	 * @param thirdPartyId 业务编号
	 * @param callbackUrl  回调地址
	 * @param remark 备注
	 * @return 提币结果
	 */
	Payout(address, currency, amount, thirdPartyId, callbackUrl, remark string) (*Result[Payout], error)

	//PayoutQuery
	/**
	 * 提币申请查看
	 * @param cid 业务编号
	 * @return 申请详细信息
	 */
	PayoutQuery(cid int64) (*Result[PayoutQuery], error)

	//AddressCreate
	/**
	 * 生成地址
	 * @param mainCoinType 主币编号
	 * @param callbackUrl 回调地址
	 * @param alias 地址别名
	 * @return 地址信息
	 */
	AddressCreate(chainId, callbackUrl, alias string) (*Result[ProjectAddress], error)
}
