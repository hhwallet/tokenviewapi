package tokenviewapi

import "fmt"

var (
	pathScbasic = "/stablecoin/getscbasic/%s"
)

type respScbasic struct {
	Balance      float64 `json:"balance"`
	Burnaddress  string  `json:"burnaddress"`
	Circulation  int64   `json:"circulation"`
	Marketvalue  float64 `json:"marketvalue"`
	Mint         float64 `json:"mint"`
	Mintaddress  string  `json:"mintaddress"`
	Network      string  `json:"network"`
	Status       int     `json:"status"`
	Tokenaddress string  `json:"tokenaddress"`
	Tokenname    string  `json:"tokenname"`
	Totalsupply  float64 `json:"totalsupply"`
}

//稳定币基本信息
//https://services.tokenview.com/vipapi/stablecoin/getscbasic/{stableCoin}?apikey={apikey}
func (t *TokenViewAPI) Getscbasic(coin string) (resp *respScbasic, err error) {
	var url = t.setUrl(pathScbasic, coin)
	fmt.Println(url)
	err = t.do2(url, &resp)
	return
}

//https://services.tokenview.com/vipapi/stablecoin/getscbasic/usdt?apikey=vEtsZFdvt7GWkKT3GEym
//https://services.tokenview.com/vipapi/stablecoin/getscbasic/usdt?apikey=tblGrv5FBZLCGU2YUJSC
