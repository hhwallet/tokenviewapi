package tokenviewapi

var (
	pathNEOAsset      = "/neo/asset/%s"
	pathNEOTokens     = "/tokens/neo/%d/%d"
	pathNEOTokenTrans = "/neo/address/tokentrans/%s/%s/%d/%d"
)

type respNEOAsset struct {
	Network     string `json:"network"`
	TokenHash   string `json:"tokenHash"`
	TransferCnt int    `json:"transferCnt"`
	HolderCnt   int    `json:"holderCnt"`
	TokenInfo   struct {
		H string `json:"h"`
		F string `json:"f"`
		S string `json:"s"`
		D string `json:"d"`
		T string `json:"t"`
		C string `json:"c"`
		I string `json:"i"`
		O string `json:"o"`
	} `json:"tokenInfo"`
}

//NEOAsset NEO资产 (代币) 信息简介
//https://services.tokenview.com/vipapi/neo/asset/{代币地址}?apikey={apikey}
func (t *TokenViewAPI) NEOAsset(addr string) (resp *respNEOAsset, err error) {
	var url = t.setUrl(pathNEOAsset, addr)
	err = t.do(url, &resp)
	return
}

type respNEOTokens struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
	Data  []struct {
		HolderCnt         int    `json:"holderCnt"`
		CirculationSupply string `json:"circulationSupply"`
		TokenInfo         struct {
			F string `json:"f"`
			S string `json:"s"`
			D string `json:"d"`
			T string `json:"t"`
			C string `json:"c"`
			I string `json:"i"`
			O string `json:"o"`
			H string `json:"h,omitempty"`
		} `json:"tokenInfo"`
	} `json:"data"`
}

//NEOTokens NEO资产(代币)列表
//https://services.tokenview.com/vipapi/tokens/neo/{页码}/{每页条数}?apikey={apikey}
func (t *TokenViewAPI) NEOTokens(page, num int) (resp *respNEOTokens, err error) {
	var url = t.setUrl(pathNEOTokens, page, num)
	err = t.do(url, &resp)
	return
}

type respNEOTokenTrans struct {
	Index         int    `json:"index"`
	BlockNo       int    `json:"block_no"`
	TokenAddr     string `json:"tokenAddr"`
	TokenSymbol   string `json:"tokenSymbol"`
	TokenDecimals string `json:"tokenDecimals"`
	Time          int    `json:"time"`
	Txid          string `json:"txid"`
	TokenInfo     struct {
		H string `json:"h"`
		F string `json:"f"`
		S string `json:"s"`
		D string `json:"d"`
		C string `json:"c"`
	} `json:"tokenInfo"`
	From          string `json:"from"`
	To            string `json:"to"`
	Value         string `json:"value"`
	Conformations int    `json:"conformations"`
}

//NEO地址的资产交易列表
//https://services.tokenview.com/vipapi/neo/address/tokentrans/{地址hash}/{资产hash}/{页码}/{每页交易条数}?apikey={apikey}
func (t *TokenViewAPI) NEOTokenTrans(addrHash, assetHash string, page, num int) (resp []*respNEOTokenTrans, err error) {
	var url = t.setUrl(pathNEOTokenTrans, addrHash, assetHash, page, num)
	err = t.do(url, &resp)
	return
}
