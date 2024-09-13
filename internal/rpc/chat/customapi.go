package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/openimsdk/tools/errs"
	"gopkg.in/yaml.v3"
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

type DjtConfig struct {
	EnterpriseUrl string `yaml:"enterpriseUrl"`
}

func iMUserVerify(djttoken string, touserid string) (*respParma, error) {
	fmt.Println("iMUserVerifyiMUserVerify")

	// 获取当前工作目录
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return nil, errs.ErrArgs.WrapMsg("Error getting working directory:", err)
	}

	fmt.Println("wd:" + wd)
	rootDir := filepath.Dir(filepath.Dir(filepath.Dir(filepath.Dir(filepath.Dir(wd)))))
	yamlFile, err := ioutil.ReadFile(rootDir + "/djt-config.yml")
	if err != nil {
		fmt.Printf("error reading YAML file: %v", err)
		return nil, errs.ErrArgs.WrapMsg("error reading YAML file: %v", err)
	}

	var djtConfig DjtConfig
	fmt.Print("读取")
	err = yaml.Unmarshal(yamlFile, &djtConfig)
	if err != nil {
		fmt.Printf("error unmarshalling YAML: %v", err)
		fmt.Print("读取失败")
		return nil, errs.ErrArgs.WrapMsg("error unmarshalling YAML: %v", err)
	}

	fmt.Print("读取EnterpriseUrl：" + djtConfig.EnterpriseUrl)

	//body参数
	reqBody := reqBodyParmas{
		Mobile: "12345678910",
	}
	reqJson, _ := json.Marshal(&reqBody)

	request, err := http.NewRequest(http.MethodGet, djtConfig.EnterpriseUrl+"/api/Enterprise/IMUserVerify", bytes.NewReader(reqJson))
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
