package tokenviewapi

var (
	pathTxPending    = "/pending/%s/%d/%d"
	pathTxPendingOne = "/pending/%s/%s"
	pathNtxPending   = "/pending/ntx//%s/%s"
	pathPendingstat  = "/pendingstat/%s"
)

type respTxPending struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
	Data  []struct {
		Type     string `json:"type"`
		Network  string `json:"network"`
		Time     int    `json:"time"`
		Txid     string `json:"txid"`
		Fee      string `json:"fee"`
		From     string `json:"from"`
		To       string `json:"to"`
		Nonce    int    `json:"nonce"`
		GasPrice int64  `json:"gasPrice"`
		GasLimit int    `json:"gasLimit"`
		Value    string `json:"value"`
		GasUsed  int    `json:"gasUsed"`
	} `json:"data"`
	Note string `json:"note"`
}

//待确认交易总体信息
//https://services.tokenview.com/vipapi/pending/{公链简称小写}/{页码}/{每页大小}?apikey={apikey}
func (t *TokenViewAPI) TxPending(coin string, page, num int) (resp *respTxPending, err error) {
	var url = t.setUrl(pathTxPending, coin, page, num)
	err = t.do(url, &resp)
	return
}

type respTxPendingOne struct {
}

//查询某笔交易是否在待确认队列里
//https://services.tokenview.com/vipapi/pending/{公链简称小写}/{交易id}?apikey={apikey}
//1. 待确认交易：交易还等待打包，还没有进入到区块链里。如果返回404，说明待确认队列里，没有这笔交易，交易已成功或失败。
//2. 如果该交易在待确认队列里，应答如下（因为还没打包，所以块高未知）：
func (t *TokenViewAPI) TxPendingOne(coin, addr string) (resp *respTxPendingOne, err error) {
	var url = t.setUrl(pathTxPendingOne, coin, addr)
	err = t.do(url, &resp)
	return
}

type respNtxPending struct {
}

//普通地址待确认交易列表 normal tx pending
//https://services.tokenview.com/vipapi/pending/ntx/{公链小写简称}/{地址}?apikey={apikey}
func (t *TokenViewAPI) NtxPending(coin, addr string) (resp *respNtxPending, err error) {
	var url = t.setUrl(pathNtxPending, coin, addr)
	err = t.do(url, &resp)
	return
}

type respPendingstat struct {
	BestGasPrice          string `json:"bestGasPrice"`
	TotalCnt              int    `json:"totalCnt"`
	MinGasPrice           string `json:"minGasPrice"`
	TotalSentValue        string `json:"totalSentValue"`
	MaxGasPrice           string `json:"maxGasPrice"`
	TotalFee              string `json:"totalFee"`
	TotalTokenTransferCnt int    `json:"totalTokenTransferCnt"`
	TotalCallTransferCnt  int    `json:"totalCallTransferCnt"`
}

//获取交易池子信息 (包括最佳手续费)
//https://services.tokenview.com/vipapi/pendingstat/{币简称小写}?apikey={apikey}
func (t *TokenViewAPI) Pendingstat(coin string) (resp *respPendingstat, err error) {
	var url = t.setUrl(pathPendingstat, coin)
	err = t.do(url, &resp)
	return
}
