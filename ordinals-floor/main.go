package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Data struct {
	Results struct {
		Floor      float64 `json:"floor"`
		Listed     float64 `json:"listedCount"`
		Marketdata struct {
			Volume24hrs float64 `json:"volume24hBtc"`
			Volume7Day  float64 `json:"volume7dBtc"`
			VolumeTotal float64 `json:"volumeTotalBtc"`
		} `json:"marketData"`
	} `json:"results"`
}

func main() {
	//input collection
	var input string
	fmt.Print("Input collection : ")
	fmt.Scanln(&input)
	//init url
	url := fmt.Sprintf("https://ordinals.market/api/nativeCollectionStats?collectionId=%s", input)

	//get resp api
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Got error ", err)
		return
	}

	//get resp body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Something went wrong ", err)
		return
	}

	//unmarshal
	var data Data
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal("Error when unmarhsal ", err)
		return
	}

	//printout
	fmt.Println("Floor : ₿", data.Results.Floor)
	fmt.Println("Listed : ", data.Results.Listed)
	fmt.Println("Volume 24 hours : ₿", data.Results.Marketdata.Volume24hrs)
	fmt.Println("Volume 7 Day : ₿", data.Results.Marketdata.Volume7Day)
	fmt.Println("Total Volume : ₿", data.Results.Marketdata.VolumeTotal)

}

//this api grabbed from https://ordinals.market/ . only for educational purpose
