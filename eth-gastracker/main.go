package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Supposes struct {
	Result struct {
		SafeGasPrice    string `json:"SafeGasPrice"`
		ProposeGasPrice string `json:"ProposeGasPrice"`
		FastGasPrice    string `json:"FastGasPrice"`
	} `json:"result"`
}

func main() {
	//init url
	url := "https://api.etherscan.io/api?module=gastracker&action=gasoracle&apikey=your api key"

	//get resp api
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Something went wrong:", err)
		return
	}
	//close
	defer resp.Body.Close()

	//get resp body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Something went wrong:", err)
		return
	}

	var response Supposes
	//unmarshal
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal("Something went wrong:", err)
		return
	}
	//print out
	fmt.Println("🐢 Low Gas :", response.Result.SafeGasPrice)
	fmt.Println("🚗 Medium Gas:", response.Result.ProposeGasPrice)
	fmt.Println("🚀 Fast Gas:", response.Result.FastGasPrice)
}
