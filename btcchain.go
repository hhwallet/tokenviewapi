package tokenviewapi

import (
	"fmt"
	"strings"
)

var (
	pathBTCBalance = "/addr/b/%s/%s"
	pathBTCUnspent = "/unspent/%s/%s/%d/%d"
	pathBTCTxlist  = "/address/%s/%s/%d/%d"
)

var allowBTCChain = map[string]string{
	"BTC": "BTC", "BCH": "BCH", "BCHSV": "BCHSV", "ADA": "ADA", "DASH": "DASH", "LTC": "LTC", "DOGE": "DOGE", "RVN": "RVN", "PIVX": "PIVX",
	"NMC": "NMC", "RDD": "RDD", "XZC": "XZC", "NRG": "NRG", "SYS": "SYS", "NEBL": "NEBL", "VTC": "VTC", "VITAE": "VITAE", "FTC": "FTC", "GIN": "GIN",
	"HC": "HC", "CRW": "CRW", "GAME": "GAME", "PART": "PART", "EMC": "EMC", "UNO": "UNO", "XSN": "XSN", "WGR": "WGR", "EMC2": "EMC2", "BCI": "BCI",
	"BLK": "BLK", "CLAM": "CLAM", "XVG": "XVG", "LUX": "LUX", "SMART": "SMART", "DCR": "DCR", "BAY": "BAY", "FLO": "FLO", "NAV": "NAV", "STRAT": "STRAT",
	"XMY": "XMY", "BCA": "BCA", "DGB": "DGB", "LCC": "LCC",
}

func (b *TokenViewAPI) chkBTCChain(coin string) error {
	if _, ok := allowBTCChain[strings.ToUpper(coin)]; !ok {
		return fmt.Errorf("不支持此币种")
	}
	return nil
}

//BTCBalance BTC类地址余额信息
//https://services.tokenview.com/vipapi/addr/b/{公链简称}/{地址}?apikey={apikey}
//{"code":1,"msg":"成功","data":"0.35832945"}
func (t *TokenViewAPI) BTCBalance(coin string, addr string) (resp string, err error) {
	var url = t.setUrl(pathBTCBalance, coin, addr)
	err = t.doBTC(coin, url, &resp)
	return
}

type respBTCUnspent struct {
	BlockNo       int    `json:"block_no"`
	OutputNo      int    `json:"output_no"`
	Tndex         string `json:"index"`
	Txid          string `json:"txid"`
	Hex           string `json:"hex"`
	Confirmations int    `json:"confirmations"`
	Value         string `json:"value"`
}

//BTCUnspent BTC类链 UTXO 列表
//https://services.tokenview.com/vipapi/unspent/{公链简称小写}/{地址}/{页码}/{每页交易条数}?apikey={apikey}
//{1 成功 [map[block_no:707226 confirmations:6 hex:76a9142ae09917462d70dcea1cd3779ccc79f0c1fbb14788ac index:6 output_no:0 txid:70cf60b49a1829420ae59af0aa7ac2027a9844a1512537020151bcf2f55579e4 value:0.01392665] map[block_no:707059 confirmations:173 hex:76a9142ae09917462d70dcea1cd3779ccc79f0c1fbb14788ac index:1914 output_no:0 txid:9ba5d0902da21cfda349cf159cf8be98273f9d8be0be6053ec0e456f76d298d3 value:0.01069044]]}
func (t *TokenViewAPI) BTCUnspent(coin string, addr string, page, num int) (resp []*respBTCUnspent, err error) {
	var url = t.setUrl(pathBTCUnspent, coin, addr, page, num)
	err = t.doBTC(coin, url, &resp)
	return
}

type respBTCTxlist struct {
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
		Inputs        []struct {
			InputNo      int    `json:"input_no"`
			Address      string `json:"address"`
			Value        string `json:"value"`
			ReceivedFrom struct {
				OutputNo int    `json:"output_no"`
				Txid     string `json:"txid"`
			} `json:"received_from"`
		} `json:"inputs"`
		Outputs []struct {
			OutputNo int    `json:"output_no"`
			Address  string `json:"address"`
			Value    string `json:"value"`
		} `json:"outputs"`
		InputCnt  int `json:"inputCnt"`
		OutputCnt int `json:"outputCnt"`
	} `json:"txs"`
}

//BTCTxlist BTC类链地址的交易列表
// https://services.tokenview.com/vipapi/address/{公链简称小写}/{地址}/{页码}/{每页交易条数}?apikey={apikey}
func (t *TokenViewAPI) BTCTxlist(coin string, addr string, page, num int) (*respBTCTxlist, error) {
	var resp []*respBTCTxlist
	var url = t.setUrl(pathBTCTxlist, coin, addr, page, num)
	err := t.doBTC(coin, url, &resp)
	if err != nil {
		return nil, err
	}
	return resp[0], nil
}
