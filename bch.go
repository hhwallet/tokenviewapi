package tokenviewapi

var pathBCHFormatAddr = "/utils/cashaddr/%s"

type respBCHFormatAddr struct {
	LegacyAddress string `json:"legacyAddress"`
	CashAddress   string `json:"cashAddress"`
}

//BCH 地址格式化
// https://services.tokenview.com/vipapi/utils/cashaddr/{新老格式地址之一}?apikey={apikey}
func (t *TokenViewAPI) BCHFormatAddr(addr string) (resp *respBCHFormatAddr, err error) {
	var url = t.setUrl(pathBCHFormatAddr, addr)
	err = t.do(url, &resp)
	return
}
