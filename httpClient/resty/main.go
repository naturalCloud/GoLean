package main

import (
	"GoLearn/testTool"
	"github.com/bytedance/sonic"
	"github.com/go-resty/resty/v2"
)

func main() {

	client := resty.New()

	response, _ := client.R().Get("https://api.bilibili.com/x/web-interface/search/default")

	var json interface{}
	sonic.Unmarshal(response.Body(), &json)
	testTool.Dd(json)
}
