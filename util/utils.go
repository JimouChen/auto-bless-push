package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type SendBody struct {
	MyServerKey   string
	MyCCServerKey string
	Title1        string
	Title2        string
	DespSuccess   string
	DespFail      string
}

func AddTimeAtEnd(text string) string {
	return text + "\n" + time.Now().String()
}

func ReadFromMdToString(mdPath string) string {
	content, err := ioutil.ReadFile(mdPath)
	if err != nil {
		fmt.Println(err.Error())
	}
	return AddTimeAtEnd(string(content))
}

func ReadDataFromJson(jsonPath string) SendBody {
	filePtr, err := os.Open(jsonPath)
	if err != nil {
		fmt.Printf("Open file failed [Err:%s]\n", err.Error())
		return SendBody{}
	}
	defer filePtr.Close()
	var info SendBody
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&info)
	if err != nil {
		fmt.Printf("Decode failed: [Err:%s]\n", err.Error())
		return SendBody{}
	}

	return info
}
