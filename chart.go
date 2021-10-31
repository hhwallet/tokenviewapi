package tokenviewapi

import (
	"encoding/json"
)

var (
	pathChart = "/chart/%s/%s?splice=%d"
)

type respChart struct {
	Code int         `json:"code"`
	Unit interface{} `json:"unit"`
	Data string      `json:"data"`
}

//曲线图接口
//https://services.tokenview.com/vipapi/chart/{公链币简称}/{图表类型}?splice=5000&apikey={apikey}
func (t *TokenViewAPI) Chart(coin string, chartType string, splice int) (resp *respChart, err error) {
	var url = t.setUrl(pathChart, coin, chartType, splice)
	var b []byte
	b, err = t.doRaw(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &resp)
	return
}
