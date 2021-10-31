package tokenviewapi

var (
	pathCoinHeight     = "/coin/latest/%s"
	pathBlockHeader    = "/block/%s/%s"
	pathBlock          = "/tx/%s/%d/%d/%d"
	pathBlocks         = "/blocks/%s/%d/%d"
	pathTxDetail       = "/tx/%s/%s"
	pathTxConfirmation = "/tx/confirmation/%s/%s"
	pathAddrMining     = "/%s/address/mining/%s/%d/%d"
)

//CoinHeight 最新块高（单条链）
//https://services.tokenview.com/vipapi/coin/latest/{币简称小写}?apikey={apikey}
func (t *TokenViewAPI) CoinHeight(coin string) (resp int64, err error) {
	var url = t.setUrl(pathCoinHeight, coin)
	err = t.do(url, &resp)
	return
}

type respBlockHeader struct {
	Type              string `json:"type"`
	Network           string `json:"network"`
	BlockNo           int    `json:"block_no"`
	Time              int    `json:"time"`
	Fee               string `json:"fee"`
	Blockhash         string `json:"blockhash"`
	Reward            string `json:"reward"`
	Size              int    `json:"size"`
	PreviousBlockhash string `json:"previous_blockhash"`
	NextBlockhash     string `json:"next_blockhash"`
	Confirmations     int    `json:"confirmations"`
	Miner             string `json:"miner"`
	Txs               []struct {
		Type          string `json:"type"`
		Network       string `json:"network"`
		BlockNo       int    `json:"block_no"`
		Height        int    `json:"height"`
		Index         int    `json:"index"`
		Time          int    `json:"time"`
		Txid          string `json:"txid"`
		Fee           string `json:"fee"`
		Confirmations int    `json:"confirmations"`
		Inputs        []struct {
			InputNo int    `json:"input_no"`
			Address string `json:"address"`
			Value   string `json:"value"`
		} `json:"inputs"`
		Outputs []struct {
			OutputNo int    `json:"output_no"`
			Address  string `json:"address"`
			Value    string `json:"value"`
		} `json:"outputs"`
		InputCnt  int `json:"inputCnt"`
		OutputCnt int `json:"outputCnt"`
	} `json:"txs"`
	SentValue        string `json:"sent_value"`
	MiningDifficulty string `json:"mining_difficulty"`
	Merkleroot       string `json:"merkleroot"`
	TxCnt            int    `json:"txCnt"`
}

//BlockHeader 区块头详情
//https://services.tokenview.com/vipapi/block/{公链简称小写}/{块高或块hash}?apikey={apikey}
func (t *TokenViewAPI) BlockHeader(coin string, s string) (resp []*respBlockHeader, err error) {
	var url = t.setUrl(pathBlockHeader, coin, s)
	err = t.do(url, &resp)
	return
}

type respBlock struct {
	Type          string `json:"type"`
	Network       string `json:"network"`
	BlockNo       int    `json:"block_no"`
	Height        int    `json:"height"`
	BlockHash     string `json:"blockHash"`
	Index         int    `json:"index"`
	Time          int    `json:"time"`
	Txid          string `json:"txid"`
	Fee           string `json:"fee"`
	Confirmations int    `json:"confirmations"`
	From          string `json:"from"`
	To            string `json:"to"`
	Nonce         int    `json:"nonce"`
	ToIsContract  int    `json:"toIsContract"`
	GasPrice      int64  `json:"gasPrice"`
	GasLimit      int    `json:"gasLimit"`
	Value         string `json:"value"`
	GasUsed       int    `json:"gasUsed"`
	TokenTransfer []struct {
		Index         int    `json:"index"`
		Token         string `json:"token"`
		TokenAddr     string `json:"tokenAddr"`
		TokenSymbol   string `json:"tokenSymbol"`
		TokenDecimals string `json:"tokenDecimals"`
		TokenInfo     struct {
			H string `json:"h"`
			F string `json:"f"`
			S string `json:"s"`
			D string `json:"d"`
		} `json:"tokenInfo"`
		From  string `json:"from"`
		To    string `json:"to"`
		Value string `json:"value"`
	} `json:"tokenTransfer"`
}

//BlockTx 区块交易列表
//https://services.tokenview.com/vipapi/{公链简称小写}/{块高}/{页码}/{每页交易条数}?apikey={apikey}
func (t *TokenViewAPI) Block(coin string, height, page, num int) (resp []*respBlock, err error) {
	var url = t.setUrl(pathBlock, coin, height, page, num)
	err = t.do(url, &resp)
	return
}

type respBlocks struct {
	Network          string `json:"network"`
	BlockNo          int    `json:"block_no"`
	Time             int    `json:"time"`
	Size             int    `json:"size"`
	TxCnt            int    `json:"txCnt"`
	SentValue        string `json:"sentValue"`
	Miner            string `json:"miner"`
	Fee              string `json:"fee"`
	Reward           string `json:"reward"`
	MiningDifficulty string `json:"miningDifficulty"`
	GasLimit         int    `json:"gasLimit"`
	GasPrice         int64  `json:"gasPrice"`
	GasUsed          int    `json:"gasUsed"`
}

//Blocks 区块列表
//https://services.tokenview.com/vipapi/blocks/{公链简称}/{页码}/{每页大小}?apikey={apikey}
func (t *TokenViewAPI) Blocks(coin string, page, num int) (resp []*respBlocks, err error) {
	var url = t.setUrl(pathBlocks, coin, page, num)
	err = t.do(url, &resp)
	return
}

type respTxDetail struct {
	Type          string `json:"type"`
	Network       string `json:"network"`
	BlockNo       int    `json:"block_no"`
	Height        int    `json:"height"`
	BlockHash     string `json:"blockHash"`
	Index         int    `json:"index"`
	Time          int    `json:"time"`
	Txid          string `json:"txid"`
	Fee           string `json:"fee"`
	Confirmations int    `json:"confirmations"`
	From          string `json:"from"`
	To            string `json:"to"`
	Nonce         int    `json:"nonce"`
	ToIsContract  int    `json:"toIsContract"`
	GasPrice      int    `json:"gasPrice"`
	GasLimit      int    `json:"gasLimit"`
	Value         string `json:"value"`
	InputData     string `json:"inputData"`
	GasUsed       int    `json:"gasUsed"`
	TraceErr      string `json:"traceErr"`
	ReceiptErr    string `json:"receiptErr"`
}

//TxDetail 某笔交易的交易详情
//https://services.tokenview.com/vipapi/tx/{公链简称小写}/{交易hash}?apikey={apikey}
func (t *TokenViewAPI) TxDetail(coin, hash string) (resp *respTxDetail, err error) {
	var url = t.setUrl(pathTxDetail, coin, hash)
	err = t.do(url, &resp)
	return
}

type respTxConfirmation struct {
	Network      string `json:"network"`
	TxId         string `json:"txId"`
	Confirmation int    `json:"confirmation"`
}

//某笔交易的确认数
//https://services.tokenview.com/vipapi/tx/confirmation/{公链简称小写}/{交易hash}?apikey={apikey}
func (t *TokenViewAPI) TxConfirmation(coin, hash string) (resp *respTxConfirmation, err error) {
	var url = t.setUrl(pathTxConfirmation, coin, hash)
	err = t.do(url, &resp)
	return
}

type respAddrMining struct {
	Type              string        `json:"type"`
	Network           string        `json:"network"`
	BlockNo           int           `json:"block_no"`
	Time              int           `json:"time"`
	Fee               string        `json:"fee"`
	Blockhash         string        `json:"blockhash"`
	Reward            string        `json:"reward"`
	Size              int           `json:"size"`
	PreviousBlockhash string        `json:"previous_blockhash"`
	Confirmations     int           `json:"confirmations"`
	Miner             string        `json:"miner"`
	Txs               []interface{} `json:"txs"`
	Merkleroot        string        `json:"merkleroot"`
	UnclesReward      string        `json:"unclesReward"`
	MixHash           string        `json:"mixHash"`
	GasLimit          int           `json:"gasLimit"`
	LowestGasPrice    int64         `json:"lowestGasPrice"`
	GasUsed           int           `json:"gasUsed"`
	Nonce             string        `json:"nonce"`
	Sha3Uncles        string        `json:"sha3Uncles"`
	ReceiptsRoot      string        `json:"receiptsRoot"`
	CallTxCnt         int           `json:"callTxCnt"`
	TokenTxCnt        int           `json:"tokenTxCnt"`
	SentValue         string        `json:"sent_value"`
	ExtraData         string        `json:"extraData"`
	MiningDifficulty  string        `json:"mining_difficulty"`
	TotalDifficulty   string        `json:"totalDifficulty"`
	TxCnt             int           `json:"txCnt"`
	Uncles            []interface{} `json:"uncles"`
	BaseFeePerGas     int64         `json:"baseFeePerGas"`
	BurntFee          string        `json:"burntFee"`
}

//ETH/ETC 等公链地址挖出的区块列表
//https://services.tokenview.com/vipapi/{公链简称小写}/address/mining/{address}/{页码}/{每页条数}?apikey={apikey}
func (t *TokenViewAPI) AddrMining(coin, addr string, page, num int) (resp []*respAddrMining, err error) {
	var url = t.setUrl(pathAddrMining, coin, addr, page, num)
	err = t.do(url, &resp)
	return
}
