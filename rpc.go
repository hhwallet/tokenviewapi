package tokenviewapi

import (
	"math/big"
)

var (
	pathOnchainWallet = "/onchainwallet/%s"
)

//发送BTC类链的 RawTransaction 上链
//https://services.tokenview.com/vipapi/onchainwallet/{币种简称小写}?apikey={apikey}
//应用场景：发送已经签完名打包好的交易上链。
//支持币种：BTC，BCH，BCHSV，BTG，BCD，LTC，ZCASH，DOGE，DASH，XMR,NEO,ONT

//var allowBcBTCChain

func (t *TokenViewAPI) Broadcast(coin string, hex string) {
	// 验证coin是否支持，
}

type reqTRXTriggersmartcontract struct {
	ContractAddress  string `json:"contract_address"`
	OwnerAddress     string `json:"owner_address"`
	FunctionSelector string `json:"function_selector"`
	Parameter        string `json:"parameter"`
	CallValue        int    `json:"call_value"`
	FeeLimit         int    `json:"fee_limit"`
	Method           string `json:"method"`
}

//TRX 创建TRC20代币交易 no test
//triggersmartcontract
func (t *TokenViewAPI) TRXTriggersmartcontract(coin string, req reqTRXTriggersmartcontract) (txid string, err error) {
	var url = t.setUrl(pathOnchainWallet, coin)
	return t.call(url, req)
}

type reqTRXCreateTx struct {
	OwnerAddress string `json:"owner_address"`
	ToAddress    string `json:"to_address"`
	Amount       int    `json:"amount"`
	Visible      bool   `json:"visible"`
	Method       string `json:"method"`
}

//TRX 广播交易上链 no test
//{"owner_address":"TGmzDFttsXBRpU8NHbx28Tb4Pq8GopZK8x","to_address":"TERJLhmbLMAJx85kzMpc91Zuan5F96FL37","amount":1,"visible":true,"method":"createtransaction"}
func (t *TokenViewAPI) TRXCreateTx(coin string, req reqTRXCreateTx) (txid string, err error) {
	var url = t.setUrl(pathOnchainWallet, coin)
	return t.call(url, req)
}

type reqTRXBc struct {
	RawData struct {
		Contract []struct {
			Parameter struct {
				Value struct {
					Amount       int    `json:"amount"`
					OwnerAddress string `json:"owner_address"`
					ToAddress    string `json:"to_address"`
				} `json:"value"`
				TypeUrl string `json:"type_url"`
			} `json:"parameter"`
			Type string `json:"type"`
		} `json:"contract"`
		RefBlockBytes string `json:"ref_block_bytes"`
		RefBlockHash  string `json:"ref_block_hash"`
		Expiration    int64  `json:"expiration"`
		Timestamp     int64  `json:"timestamp"`
	} `json:"raw_data"`
	RawDataHex string   `json:"raw_data_hex"`
	Signature  []string `json:"signature"`
	TxID       string   `json:"txID"`
	Visible    bool     `json:"visible"`
	Method     string   `json:"method"`
}

// BroadcastTRX TRX 广播交易上链 no test
//{"raw_data":{"contract":[{"parameter":{"value":{"amount":1,"owner_address":"TTRy1wgrn3gxeZybWNdKuPYU64kBLQ6Zmd","to_address":"TGmzDFttsXBRpU8NHbx28Tb4Pq8GopZK8x"},"type_url":"type.googleapis.com/protocol.TransferContract"},"type":"TransferContract"}],"ref_block_bytes":"2ba9","ref_block_hash":"efa970b152152a86","expiration":1606052703000,"timestamp":1606052647030},"raw_data_hex":"0a022ba92208efa970b152152a864098d6cd81df2e5a65080112610a2d747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e5472616e73666572436f6e747261637412300a1541bf89b206a6f347baf1de420f0356dccf62f2221f1215414aa9a0ed45321706f2618384a399e3fee73f13fc180170f6a0ca81df2e","signature":["7584150f82bd13ece4861e7fdb67cc4efe20da30e8467c3026db8a36951dbadca92b86a0d23a43b251bad138925f8c6025aedb9eed2c9785af1832f3d6b2b0dd01"],"txID":"923db55dd6bf012901b715debda10e2fb60fc1ff618747464da6c19720aa6188","visible":true,"method":"broadcasttransaction"}
func (t *TokenViewAPI) TRXBroadcast(coin string, req reqTRXBc) (txid string, err error) {
	var url = t.setUrl(pathOnchainWallet, coin)
	return t.call(url, req)
}

//发送BTC类链的 RawTransaction 上链
func (t *TokenViewAPI) BTCBroadcast(coin string, hex string) (txid string, err error) {
	var url = t.setUrl(pathOnchainWallet, coin)
	params := t.rpcParams("sendrawtransaction", hex)
	return t.call(url, params)
}

//发送ETH/ETC 的 RawTransaction 上链
func (t *TokenViewAPI) ETHBroadcast(coin string, hex string) (txid string, err error) {
	var url = t.setUrl(pathOnchainWallet, coin)
	params := t.rpcParams("eth_sendRawTransaction", hex)
	return t.call(url, params)
}

//返回错误：not allow method
func (t *TokenViewAPI) GasPrice(coin string) (big.Int, error) {
	var url = t.setUrl(pathOnchainWallet, coin)
	params := t.rpcParams("eth_gasPrice")
	resp, err := t.call(url, params)
	if err != nil {
		return big.Int{}, err
	}
	return ParseBigInt(resp)
}

//ETHNonce ETH/ETC 获取地址的 Nonce
func (t *TokenViewAPI) ETHNonce(coin string, address, block string) (int, error) {
	var url = t.setUrl(pathOnchainWallet, coin)
	params := t.rpcParams("eth_getTransactionCount", address, block)
	resp, err := t.call(url, params)
	if err != nil {
		return 0, err
	}
	return ParseInt(resp)
}

type ReqTx struct {
	From     string
	To       string
	Gas      int
	GasPrice *big.Int
	Value    *big.Int
	Data     string
	Nonce    int
}

//ETHEstimateGas ETH/ETC 预估Gas花费量
func (t *TokenViewAPI) ETHEstimateGas(coin string, req ReqTx) (int, error) {
	var url = t.setUrl(pathOnchainWallet, coin)
	params := t.rpcParams("eth_estimateGas", req)
	resp, err := t.call(url, params)
	if err != nil {
		return 0, err
	}
	return ParseInt(resp)
}

func (t *TokenViewAPI) rpcParams(method string, params ...interface{}) map[string]interface{} {
	p := make(map[string]interface{})
	p["jsonrpc"] = "2.0"
	p["id"] = "viewtoken"
	p["method"] = method
	p["params"] = params
	return p
}
