package chat

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// 请求体 body参数 结构体
type reqBodyParmas struct {
	Mobile string `json:"mobile"`
}

// 根据具体返回的值填写
type respParma struct {
	Status  int32  `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func HttpDo(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func iMUserVerify(djttoken string, touserid string) (*respParma, error) {

	//body参数
	reqBody := reqBodyParmas{
		Mobile: "12345678910",
	}
	reqJson, _ := json.Marshal(&reqBody)

	request, err := http.NewRequest(http.MethodGet, "http://enterprise-admin-dev.51djt.net/api/Enterprise/IMUserVerify", bytes.NewReader(reqJson))
	if err != nil {
		return nil, err
	}

	//url参数
	parma := request.URL.Query()
	parma.Add("touserid", touserid)

	request.URL.RawQuery = parma.Encode()

	//header参数
	request.Header.Set("Authorization", "Bearer "+djttoken)

	respJson, err := HttpDo(request)
	if err != nil {
		return nil, err
	}

	var resp respParma
	_ = json.Unmarshal(respJson, &resp)

	return &resp, nil
}
