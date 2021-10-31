package tokenviewapi

/**
稳定币 todo
*/

import (
	"fmt"
	"strings"
)

var (
	pathETHBalance                = "/addr/b/%s/%s"
	pathETHAddressInfo            = "/%s/address/%s"
	pathETHTxlist                 = "/address/%s/%s/%d/%d"
	pathETHTokens                 = "/%s/tokens/%d/%d"
	pathETHTokenInfo              = "/%s/token/%s"
	pathETHTokenMaxSupply         = "/token/maxsupply/%s/%s"
	pathETHTokenTotalSupply       = "/token/totalsupply/%s/%s"
	pathETHTokenCirculatingSupply = "/token/circulatingsupply/%s/%s"
	pathETHTokenTrans             = "/%s/token/tokentrans/%s/%d/%d"
	pathETHTokenAddrBal           = "/%s/address/tokenbalance/%s"
	pathETHTokenAddrTrans         = "/%s/address/tokentrans/%s/%s/%d/%d"
	pathETHCallTrans              = "/%s/address/calltrans/%s/%d/%d"
	pathETHRunlog                 = "/%s/tx/runlog/%d/%d"
)
var allowETHChain = map[string]string{
	"ETH":  "ETH",
	"ETC":  "ETC",
	"TRX":  "TRX",
	"NEO":  "NEO",
	"ONT":  "ONT",
	"IOST": "IOST",
	"UGAS": "UGAS",
	"ATOM": "ATOM",
	"WAN":  "WAN",
	"PI":   "PI",
}

func (b *TokenViewAPI) chkETHChain(coin string) error {
	if _, ok := allowETHChain[strings.ToUpper(coin)]; !ok {
		return fmt.Errorf("不支持此币种")
	}
	return nil
}

//ETHBalance ETH类链地址余额
//https://services.tokenview.com/vipapi/addr/b/{公链简称小写}/{地址}?apikey={apikey}

func (t *TokenViewAPI) ETHBalance(coin string, addr string) (resp string, err error) {
	var url = t.setUrl(pathETHBalance, coin, addr)
	err = t.doETH(coin, url, &resp)
	return
}

type respETHAddressInfo struct {
	Type          string `json:"type"`
	Network       string `json:"network"`
	Hash          string `json:"hash"`
	Spend         string `json:"spend"`
	Receive       string `json:"receive"`
	Balance       string `json:"balance"`
	NormalTxCount int    `json:"normalTxCount"`
	Txs           []struct {
		Type          string `json:"type"`
		Network       string `json:"network"`
		BlockNo       int    `json:"block_no"`
		Height        int    `json:"height"`
		Index         int    `json:"index"`
		Time          int    `json:"time"`
		Txid          string `json:"txid"`
		Fee           string `json:"fee"`
		Confirmations int    `json:"confirmations"`
		From          string `json:"from"`
		To            string `json:"to"`
		Nonce         int    `json:"nonce"`
		GasPrice      int64  `json:"gasPrice"`
		GasLimit      int    `json:"gasLimit"`
		Value         string `json:"value"`
		GasUsed       int    `json:"gasUsed"`
	} `json:"txs"`
	AddressType string `json:"addressType"`
	Nonce       int    `json:"nonce"`
}

//ETHAddressInfo ETH类链地址详情
//https://services.tokenview.com/vipapi/{公链简称}/address/{地址}?apikey={apikey}
func (t *TokenViewAPI) ETHAddressInfo(coin, addr string) (resp *respETHAddressInfo, err error) {
	var url = t.setUrl(pathETHAddressInfo, coin, addr)
	err = t.doETH(coin, url, &resp)
	return
}

type respETHTxlist struct {
	Type          string `json:"type"`
	Network       string `json:"network"`
	Hash          string `json:"hash"`
	TxCount       int    `json:"txCount"`
	Spend         string `json:"spend"`
	Receive       string `json:"receive"`
	NormalTxCount int    `json:"normalTxCount"`
	Txs           []struct {
		Type          string `json:"type"`
		Network       string `json:"network"`
		BlockNo       int    `json:"block_no"`
		Height        int    `json:"height"`
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
			From         string `json:"from"`
			To           string `json:"to"`
			Value        string `json:"value"`
			ToIsContract int    `json:"toIsContract,omitempty"`
		} `json:"tokenTransfer,omitempty"`
		TraceErr   string `json:"traceErr,omitempty"`
		ReceiptErr string `json:"receiptErr,omitempty"`
	} `json:"txs"`
}

//ETHTxlist ETH类链地址交易列表
//https://services.tokenview.com/vipapi/{公链简称小写}/address/normal/{地址}/{页码}/{每页交易条数}?apikey={apikey}
func (t *TokenViewAPI) ETHTxlist(coin string, addr string, page, num int) (*respETHTxlist, error) {
	var url = t.setUrl(pathETHTxlist, coin, addr, page, num)
	var resp []*respETHTxlist
	err := t.doETH(coin, url, &resp)
	if err != nil {
		return nil, err
	}
	return resp[0], nil
}

type respETHTokens struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
	Data  []struct {
		H  string `json:"h"`
		P  string `json:"p"`
		F  string `json:"f"`
		S  string `json:"s"`
		D  string `json:"d"`
		T  string `json:"t"`
		Ic string `json:"ic"`
	} `json:"data"`
	Note interface{} `json:"note"`
}

// ETHTokens ETH类链的代币列表
//https://services.tokenview.com/vipapi/{公链简称小写}/tokens/{页码}/{每页条数}?apikey={apikey}
func (t *TokenViewAPI) ETHTokens(coin string, page, num int) (resp *respETHTokens, err error) {
	var url = t.setUrl(pathETHTokens, coin, page, num)
	err = t.doETH(coin, url, &resp)
	return

}

type respETHTokenInfo struct {
	Network     string `json:"network"`
	TokenHash   string `json:"tokenHash"`
	TransferCnt int    `json:"transferCnt"`
	HolderCnt   int    `json:"holderCnt"`
	TokenInfo   struct {
		H string `json:"h"`
		P string `json:"p"`
		F string `json:"f"`
		S string `json:"s"`
		D string `json:"d"`
		T string `json:"t"`
	} `json:"tokenInfo"`
}

//ETHTokenInfo 代币信息简介
//https://services.tokenview.com/vipapi/{公链简称}/token/{代币地址}?apikey={apikey}
func (t *TokenViewAPI) ETHTokenInfo(coin, addr string) (resp *respETHTokenInfo, err error) {
	var url = t.setUrl(pathETHTokenInfo, coin, addr)
	err = t.doETH(coin, url, &resp)
	return
}

//ETHTokenMaxSupply 代币最大发行量
//https://services.tokenview.com/vipapi/token/maxsupply/{公链简称}/{代币地址}?apikey={apikey}
func (t *TokenViewAPI) ETHTokenMaxSupply(coin, addr string) (string, error) {
	var url = t.setUrl(pathETHTokenMaxSupply, coin, addr)
	b, err := t.doRawETH(coin, url)
	if err != nil {
		return "", err
	}
	return string(b), err
}

//ETHTokenTotalSupply 代币总共发行量
//https://services.tokenview.com/vipapi/token/totalsupply/{公链简称}/{代币地址}?apikey={apikey}
func (t *TokenViewAPI) ETHTokenTotalSupply(coin, addr string) (string, error) {
	var url = t.setUrl(pathETHTokenTotalSupply, coin, addr)
	b, err := t.doRawETH(coin, url)
	if err != nil {
		return "", err
	}
	return string(b), err
}

//ETHTokenCirculatingSupply 代币流通量
//https://services.tokenview.com/vipapi/token/circulatingsupply/{公链小写简称}/{代币地址或名称}?apikey={apikey}
func (t *TokenViewAPI) ETHTokenCirculatingSupply(coin, addr string) (string, error) {
	var url = t.setUrl(pathETHTokenCirculatingSupply, coin, addr)
	b, err := t.doRawETH(coin, url)
	if err != nil {
		return "", err
	}
	return string(b), err
}

type respETHTokentrans struct {
	Index         int    `json:"index"`
	BlockNo       int    `json:"block_no"`
	Token         string `json:"token"`
	Time          int    `json:"time"`
	Txid          string `json:"txid"`
	From          string `json:"from"`
	To            string `json:"to"`
	Value         string `json:"value"`
	Conformations int    `json:"conformations"`
}

//ETHTokenTrans 代币（本身）交易列表
//https://services.tokenview.com/vipapi/{公链简称小写}/token/tokentrans/{代币地址}/{页码}/{每页交易条数}?apikey={apikey}
func (t *TokenViewAPI) ETHTokenTrans(coin string, tokenAddr string, page, num int) (resp []*respETHTokentrans, err error) {
	var url = t.setUrl(pathETHTokenTrans, coin, tokenAddr, page, num)
	err = t.doETH(coin, url, &resp)
	return
}

type respETHTokenAddrBal struct {
	Network   string `json:"network"`
	Hash      string `json:"hash"`
	TokenInfo struct {
		H string `json:"h"`
		F string `json:"f"`
		S string `json:"s"`
		D string `json:"d"`
	} `json:"tokenInfo"`
	TransferCnt int    `json:"transferCnt"`
	Balance     string `json:"balance"`
}

//ETHTokenAddrBal 某个地址的代币余额列表
//https://services.tokenview.com/{公链简称小写}/address/tokenbalance/{持币地址}?apikey={apikey}
func (t *TokenViewAPI) ETHTokenAddrBal(coin string, addr string) (resp []*respETHTokenAddrBal, err error) {
	var url = t.setUrl(pathETHTokenAddrBal, coin, addr)
	err = t.doETH(coin, url, &resp)
	return
}

type respETHTokenAddrTrans struct {
	Index         int    `json:"index"`
	BlockNo       int    `json:"block_no"`
	Token         string `json:"token"`
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
	} `json:"tokenInfo"`
	From          string `json:"from"`
	To            string `json:"to"`
	Value         string `json:"value"`
	Conformations int    `json:"conformations"`
}

//ETHTokenAddrTrans 某个地址的代币交易列表
//https://services.tokenview.com/vipapi/{公链简称小写}/address/tokentrans/{地址}/{代币地址}/{页码}/{每页交易条数}?apikey={apikey}
func (t *TokenViewAPI) ETHTokenAddrTrans(coin string, addr, tokenAddr string, page, num int) (resp []*respETHTokenAddrTrans, err error) {
	var url = t.setUrl(pathETHTokenAddrTrans, coin, addr, tokenAddr, page, num)
	err = t.doETH(coin, url, &resp)
	return
}

type respETHCallTrans struct {
	BlockNo        int    `json:"block_no"`
	Txid           string `json:"txid"`
	Time           int    `json:"time"`
	Index          int    `json:"index"`
	FromIsContract int    `json:"fromIsContract"`
	From           string `json:"from"`
	To             string `json:"to"`
	Value          string `json:"value"`
	GasUsed        int    `json:"gasUsed"`
	GasLimit       int    `json:"gasLimit"`
	Conformations  int    `json:"conformations"`
}

//ETHCallTrans 某个地址的合约调用转账列表
//https://services.tokenview.com/vipapi/{公链简称小写}/address/calltrans/{address}/{页码}/{每页条数}?apikey={apikey}
func (t *TokenViewAPI) ETHCallTrans(coin string, addr string, page, num int) (resp []*respETHCallTrans, err error) {
	var url = t.setUrl(pathETHCallTrans, coin, addr, page, num)
	err = t.doETH(coin, url, &resp)
	return
}

type respETHRunlog struct {
	BlockNo int `json:"block_no"`
	Trace   []struct {
		Result struct {
			Output  string `json:"output"`
			GasUsed string `json:"gasUsed"`
		} `json:"result"`
		TraceAddress []interface{} `json:"traceAddress"`
		Action       struct {
			Input    string `json:"input"`
			Gas      string `json:"gas"`
			From     string `json:"from"`
			To       string `json:"to"`
			Value    string `json:"value"`
			CallType string `json:"callType"`
		} `json:"action"`
		Type      string `json:"type"`
		Subtraces int    `json:"subtraces"`
	} `json:"trace"`
	Txid    string `json:"txid"`
	Index   int    `json:"index"`
	Receipt struct {
		LogsBloom         string `json:"logsBloom"`
		GasUsed           string `json:"gasUsed"`
		CumulativeGasUsed string `json:"cumulativeGasUsed"`
		Logs              []struct {
			Address             string   `json:"address"`
			TransactionLogIndex string   `json:"transactionLogIndex"`
			LogIndex            string   `json:"logIndex"`
			Data                string   `json:"data"`
			Removed             bool     `json:"removed"`
			Topics              []string `json:"topics"`
			Type                string   `json:"type"`
		} `json:"logs"`
		Status string `json:"status"`
	} `json:"receipt"`
}

//ETHRunlog ETH 交易的日志详情
//https://services.tokenview.com/vipapi/{公链简称小写}/tx/runlog/{交易所在区块高度}/{交易在区块里的位置}?apikey={apikey}
func (t *TokenViewAPI) ETHRunlog(coin string, height, position int) (resp *respETHRunlog, err error) {
	var url = t.setUrl(pathETHRunlog, coin, height, position)
	err = t.doETH(coin, url, &resp)
	return
}
