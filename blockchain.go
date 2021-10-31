package tokenviewapi

var (
	pathBlockHeight    = "/block/latest/height"
	pathBlockInfo      = "/coin/abstract/%s"
	pathCoinMaxsupply  = "/coin/maxsupply/%s"
	pathRichrange      = "/address/richrange/%s/%d/%d"
	pathTokenRichrange = "/address/richrange/%s_%s/%d/%d"
	pathHashtype       = "/hashtype/%s"
)

//BlockHeight 公链列表
//https://services.tokenview.com/vipapi/block/latest/height?apikey={apikey}
func (t *TokenViewAPI) BlockHeight() (resp map[string]string, err error) {
	var url = t.setUrl(pathBlockHeight)
	err = t.do(url, &resp)
	return
}

//BlockInfo 公链基本信息
//https://services.tokenview.com/vipapi/coin/abstract/{公链简称小写}?apikey={apikey}
func (t *TokenViewAPI) BlockInfo(coin string) (resp map[string]string, err error) {
	var url = t.setUrl(pathBlockInfo, coin)
	err = t.do(url, &resp)
	return
}

//CoinMaxsupply 最大发行量
//https://services.tokenview.com/vipapi/coin/maxsupply/{币简称小写}?apikey={apikey}
func (t *TokenViewAPI) CoinMaxsupply(coin string) (resp string, err error) {
	var url = t.setUrl(pathCoinMaxsupply, coin)
	b, err := t.doRaw(url)
	if err != nil {
		return "", err
	}
	return string(b), err
}

type respRichrange struct {
	Addr    string  `json:"addr"`
	Balance float64 `json:"balance"`
	TxCnt   int     `json:"txCnt"`
}

//Richrange 公链富豪榜Top200列表
//https://services.tokenview.com/vipapi/address/richrange/{公链简称小写}/1/200?apikey={apikey}
func (t *TokenViewAPI) Richrange(coin string, page, num int) (resp []*respRichrange, err error) {
	var url = t.setUrl(pathRichrange, coin, page, num)
	err = t.do(url, &resp)
	return
}

type respTokenRichrange struct {
	Addr    string  `json:"addr"`
	Balance float64 `json:"balance"`
	TxCnt   int     `json:"txCnt"`
}

//TokenRichrange 代币富豪榜Top200列表
//https://services.tokenview.com/address/richrange/{公链简称小写}_{代币地址}/1/200?apikey={apikey}
func (t *TokenViewAPI) TokenRichrange(coin, addr string, page, num int) (resp []*respTokenRichrange, err error) {
	var url = t.setUrl(pathTokenRichrange, coin, addr, page, num)
	err = t.do(url, &resp)
	return
}

//查询区块链hash的类型
//https://services.tokenview.com/vipapi/hashtype/{任意hash字符串}?apikey={apikey}
func (t *TokenViewAPI) Hashtype(hash string) (resp map[string]string, err error) {
	var url = t.setUrl(pathHashtype, hash)
	err = t.do(url, &resp)
	return
}
