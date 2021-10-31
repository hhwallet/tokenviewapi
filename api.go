package tokenviewapi

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"strings"
)

const DefaultAPIHost = "https://services.tokenview.com/vipapi"

type TokenViewAPI struct {
	Key  string
	Host string
}

type response struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}

type response2 struct {
	Code int             `json:"code"`
	Data json.RawMessage `json:"data"`
}

type respBroadcast struct {
	Result  string    `json:"result"`
	Error   respError `json:"error"`
	Jsonrpc string    `json:"jsonrpc"`
	Id      string    `json:"id"`
}

type respError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func New(key string, host string) *TokenViewAPI {
	if host == "" {
		host = DefaultAPIHost
	}

	return &TokenViewAPI{
		key, host,
	}
}

func (t *TokenViewAPI) Balance(coin string, address string) {
	return
}

func (t *TokenViewAPI) getRaw(url string) ([]byte, error) {
	var b []byte
	resp, err := grequests.Get(url, nil)

	if err != nil {
		return b, fmt.Errorf("请求失败：%s", err.Error())
	}
	if resp.StatusCode != 200 {
		return b, fmt.Errorf("状态码出错：%d", resp.StatusCode)
	}
	//fmt.Println("resp.String()", resp.String())
	return resp.Bytes(), nil

}

func (t *TokenViewAPI) get(url string) (b json.RawMessage, err error) {
	var r response
	b, err = t.getRaw(url)
	if err != nil {
		return
	}

	if err = json.Unmarshal(b, &r); err != nil {
		return nil, fmt.Errorf("JSON解析出错：,%s", err.Error())
	}
	if r.Code != 1 {
		return nil, fmt.Errorf(r.Msg)
	}
	return r.Data, nil

}

func (t *TokenViewAPI) get2(url string) (b json.RawMessage, err error) {
	var r response2
	b, err = t.getRaw(url)
	if err != nil {
		return
	}

	if err = json.Unmarshal(b, &r); err != nil {
		return nil, fmt.Errorf("JSON解析出错：,%s", err.Error())
	}
	if r.Code != 1 {
		return nil, fmt.Errorf("查询出错")
	}
	return r.Data, nil

}

func (t *TokenViewAPI) setUrl(path string, p ...interface{}) string {
	l := t.Host + path
	l = fmt.Sprintf(l, p...)
	if strings.Index(l, "?") == -1 {
		l = l + "?apikey=" + t.Key
	} else {
		l = l + "&apikey=" + t.Key
	}
	return l
}

func (t *TokenViewAPI) doBTC(coin string, url string, resp interface{}) (err error) {
	if err = t.chkBTCChain(coin); err != nil {
		return
	}
	return t.do(url, resp)
}

func (t *TokenViewAPI) doETH(coin string, url string, resp interface{}) (err error) {
	if err = t.chkETHChain(coin); err != nil {
		return
	}
	return t.do(url, resp)
}
func (t *TokenViewAPI) do(url string, resp interface{}) (err error) {
	var res json.RawMessage
	if res, err = t.get(url); err != nil {
		return
	} else {
		err = json.Unmarshal(res, resp)
		return
	}
}

func (t *TokenViewAPI) do2(url string, resp interface{}) (err error) {
	var res json.RawMessage
	if res, err = t.get2(url); err != nil {
		return
	} else {
		err = json.Unmarshal(res, resp)
		return
	}
}

func (t *TokenViewAPI) doRawBTC(coin string, url string) ([]byte, error) {
	if err := t.chkBTCChain(coin); err != nil {
		return nil, err
	}
	return t.doRaw(url)
}

func (t *TokenViewAPI) doRawETH(coin string, url string) ([]byte, error) {
	if err := t.chkETHChain(coin); err != nil {
		return nil, err
	}
	return t.doRaw(url)
}

func (t *TokenViewAPI) doRaw(url string) ([]byte, error) {
	return t.getRaw(url)
}

//curl --location --request POST 'https://wallet.tokenview.com/onchainwallet/btc' \
//--header 'Content-Type: application/json' \
//--data-raw '{"jsonrpc": "1.0", "id":"viewtoken", "method": "sendrawtransaction", "params": ["0100000001e5269c07fd5f609fe0cc479e365113f3054a8ebfc89ebd6d188defd3a2cc3c75010000006b483045022100c11aa61cebf9f4a337b07ef34b018a722a83a44f5f1db94d149cf7086f93aafa02200e2df5dfa6025a7292e5fc07a2508c98f4afe4bf59831f4c5fe1513728c5e80301210288521c776e00498c268215d12288b2b1ec87ba9b9bcee570df8b3da038d01f34ffffffff0222020000000000001976a9147e3f596889eb2c4666ea0823f05b0b7e0796192388ac79230000000000001976a9143a05fb822a4954af7d868d133963856fa3b5d7de88ac00000000"] }'
func (t *TokenViewAPI) call(url string, params interface{}) (string, error) {

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	resp, err := grequests.Post(url, &grequests.RequestOptions{
		Headers: headers,
		JSON:    params,
	})
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf(resp.String())
	}
	respBc := respBroadcast{}
	err = json.Unmarshal(resp.Bytes(), &respBc)
	if err != nil {
		return "", err
	}
	if respBc.Result == "" {
		return "", fmt.Errorf(respBc.Error.Message)
	} else {
		return respBc.Result, nil
	}

}
