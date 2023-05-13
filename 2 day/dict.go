package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// 请求头结构体
type DictRequest struct {
	// 翻译类型
	TransType string `json:"trans_type"`
	// 翻译文本
	Source string `json:"source"`
	// 用户id
	UserID string `json:"user_id"`
}

// 响应数据文本，少数参数有用
type DictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		// 翻译结果
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []interface{} `json:"antonym"`
		// 可使用词组
		WqxExample [][]string `json:"wqx_example"`
		// 翻译文本
		Entry   string        `json:"entry"`
		Type    string        `json:"type"`
		Related []interface{} `json:"related"`
		Source  string        `json:"source"`
	} `json:"dictionary"`
}

func query(word string) {
	client := &http.Client{}
	// 设置请求参数
	request := DictRequest{TransType: "en2zh", Source: word}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buf)

	// 设置参数数据流
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	// 请求头
	req.Header.Set("authority", "api.interpreter.caiyunai.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("app-name", "xy")
	req.Header.Set("content-type", "application/json;charset=UTF-8")
	req.Header.Set("device-id", "f1de93819e3bb9f68a199a51c6ee2efb")
	req.Header.Set("origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("os-type", "web")
	req.Header.Set("os-version", "")
	req.Header.Set("referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("sec-ch-ua", `"Microsoft Edge";v="113", "Chromium";v="113", "Not-A.Brand";v="24"`)
	req.Header.Set("sec-ch-ua-mobile", "?1")
	req.Header.Set("sec-ch-ua-platform", `"Android"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Mobile Safari/537.36 Edg/113.0.1774.35")
	req.Header.Set("x-authorization", "token:qgemv4jr1y38jyq6vhvi")
	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	// 关闭请求流
	defer resp.Body.Close()
	// 读取响应数据
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// 防止请求出错
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}

	var dictResponse DictResponse
	// 将响应数据转化为字符串
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(word, "UK:", dictResponse.Dictionary.Prons.En, "US:", dictResponse.Dictionary.Prons.EnUs)
	// 循环查找响应数据中的翻译结果
	for _, item := range dictResponse.Dictionary.Explanations {
		fmt.Println(item)
	}
}

func main() {
	// 运行代码：go run dict.go hello
	// hello 即为要翻译的文本
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD example: simpleDict hello`)
		os.Exit(1)
	}
	word := os.Args[1]
	query(word)
}
