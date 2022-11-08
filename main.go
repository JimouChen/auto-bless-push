package main

import (
	"bless-push/util"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func postPushBless(serverKey string, title string, desp string) bool {
	//postUrl := "https://sctapi.ftqq.com/" + serverKey + ".send?title=" + title + "&desp=" + desp
	postUrl := "https://sctapi.ftqq.com/" + serverKey + ".send"

	data := make(url.Values)
	data.Add("title", title)
	data.Add("desp", desp)

	request, err := http.NewRequest(http.MethodPost, postUrl, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	fmt.Println(response.Status)
	return true
}

func main() {
	markdownPath := "./text/bless.md"
	configPath := "./config.json"
	config := utils.ReadDataFromJson(configPath)

	if postPushBless(config.MyCCServerKey, config.Title1, utils.ReadFromMdToString(markdownPath)) {
		time.Sleep(time.Second * 2)
		postPushBless(config.MyServerKey, config.Title2, utils.AddTimeAtEnd(config.DespSuccess))
	} else {
		time.Sleep(time.Second * 2)
		postPushBless(config.MyServerKey, config.Title2, utils.AddTimeAtEnd(config.DespFail))
	}
}
