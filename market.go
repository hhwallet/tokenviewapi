package tokenviewapi

import (
	"encoding/json"
)

var (
	pathMarketInfo   = "/coin/marketInfo/%s"
	pathMarketCap    = "/market/marketCap?page=%d&size=%d"
	pathMarketEx     = "/market/exchange"
	pathCoinCirSup   = "/coin/circulatingsupply/%s"
	pathCoinTotalSup = "/coin/totalsupply/%s"
)

type respMarketInfo struct {
	Id                string `json:"_id"`
	ChangeUsd1H       string `json:"changeUsd1h"`
	ChangeUsd7D       string `json:"changeUsd7d"`
	ChangeUsd24H      string `json:"changeUsd24h"`
	PriceUsd          string `json:"priceUsd"`
	PriceBtc          string `json:"priceBtc"`
	VolumeUsd         string `json:"volumeUsd"`
	Name              string `json:"name"`
	Erc20Addr         string `json:"erc20Addr"`
	MarketCapUsd      string `json:"marketCapUsd"`
	Symbol            string `json:"symbol"`
	Rank              string `json:"rank"`
	UniqueId          string `json:"uniqueId"`
	CirculatingSupply string `json:"circulatingSupply"`
	TotalSupply       string `json:"totalSupply"`
	High24H           string `json:"high24h"`
	Low24H            string `json:"low24h"`
	PartialObject     bool   `json:"partialObject"`
}

//数字货币行情
//https://services.tokenview.com/vipapi/coin/marketInfo/{数字货币名称}?apikey={apikey}
func (t *TokenViewAPI) MarketInfo(name string) (resp *respMarketInfo, err error) {
	var url = t.setUrl(pathMarketInfo, name)
	err = t.do(url, &resp)
	return
}

type respMarketCap struct {
	Total         int `json:"total"`
	Page          int `json:"page"`
	Size          int `json:"size"`
	MarketCapData []struct {
		Rank                         int     `json:"rank"`
		Name                         string  `json:"name"`
		Nickname                     string  `json:"nickname"`
		MarketCapBtc                 int64   `json:"marketCapBtc"`
		ChangeUsd1H                  float64 `json:"changeUsd1h"`
		Erc20Addr                    string  `json:"erc20Addr"`
		CirculatingSupplyBtc         float64 `json:"circulatingSupplyBtc"`
		ChangeBtc7D                  int     `json:"changeBtc7d"`
		MarketCapUsd                 int64   `json:"marketCapUsd"`
		VolumeUsd                    int64   `json:"volumeUsd"`
		ChangUsd24H                  float64 `json:"changUsd24h"`
		ChangeUsd24H                 float64 `json:"changeUsd24h"`
		VolumeBtc                    int     `json:"volumeBtc"`
		UniqueId                     string  `json:"uniqueId"`
		ChangUsd7D                   float64 `json:"changUsd7d"`
		ChangeBtc24H                 int     `json:"changeBtc24h"`
		PriceBtc                     float64 `json:"priceBtc"`
		ChangeBtc1H                  int     `json:"changeBtc1h"`
		PriceUsd                     float64 `json:"priceUsd"`
		Symbol                       string  `json:"symbol"`
		MarketCapChangePercentage24H float64 `json:"market_cap_change_percentage_24h"`
	} `json:"marketCapData"`
}

//数字货币行情 (批量）
//https://services.tokenview.com/vipapi/market/marketCap?page=0&size=1000&apikey={apikey}
func (t *TokenViewAPI) MarketCap(page, num int) (resp *respMarketCap, err error) {
	var url = t.setUrl(pathMarketCap, page, num)
	b, err := t.doRaw(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &resp)
	return
}

//数字货币汇率
//https://services.tokenview.com/vipapi/market/exchange?apikey={apikey}
func (t *TokenViewAPI) MarketEx() (resp map[string]float64, err error) {
	var url = t.setUrl(pathMarketEx)
	b, err := t.doRaw(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &resp)
	return
}

//CoinCirSup 数字货币流通量
//https://services.tokenview.com/vipapi/coin/circulatingsupply/{数字货币简称小写}?apikey={apikey}
func (t *TokenViewAPI) CoinCirSup(coin string) (resp string, err error) {
	var url = t.setUrl(pathCoinCirSup, coin)
	b, err := t.doRaw(url)
	if err != nil {
		return "", err
	}
	return string(b), err
}

//数字货币总共发行量
//https://services.tokenview.com/vipapi/coin/totalsupply/{数字货币简称小写}?apikey={apikey}
func (t *TokenViewAPI) CoinTotalSup(coin string) (resp string, err error) {
	var url = t.setUrl(pathCoinTotalSup, coin)
	b, err := t.doRaw(url)
	if err != nil {
		return "", err
	}
	return string(b), err
}
