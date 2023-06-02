package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type items struct {
	TotalVolume         string `json:"totalVolume"`
	Owners              string `json:"owners"`
	Supply              string `json:"supply"`
	FloorPrice          string `json:"floorPrice"`
	TotalListed         string `json:"totalListed"`
	PendingTransactions string `json:"pendingTransactions"`
}

// format total volume ny
func formatTotalVolume(totalVolume string) string {
	// Mengkonversi totalVolume ke float64
	value, _ := strconv.ParseFloat(totalVolume, 64)

	// Mengubah angka menjadi string dengan 4 angka di belakang koma
	formattedValue := strconv.FormatFloat(value/1e8, 'f', 4, 64)

	return formattedValue
}

// format total floor price
func formatFloorPrice(floorPrice string) string {
	// Menghilangkan karakter non-digit dari floorPrice
	floorPrice = strings.ReplaceAll(floorPrice, ".", "")

	// Mengkonversi floorPrice ke float64
	value, _ := strconv.ParseFloat(floorPrice, 64)

	// Mengubah angka menjadi string dengan 4 angka di belakang koma
	formattedValue := strconv.FormatFloat(value/1e8, 'f', 4, 64)

	return formattedValue
}

func main() {

	fmt.Print("Input collection : ")
	var koleksi string
	fmt.Scanln(&koleksi)
	//init url dan input koleksi
	url := fmt.Sprintf("https://api-mainnet.magiceden.io/v2/ord/btc/stat?collectionSymbol=%s", koleksi)

	// get response dari api dan handle error
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Fetching error ", err)
		os.Exit(1)
	}

	//read respon api
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Something went wrong", err)
		return
	}

	var item items
	//json response
	err = json.Unmarshal(data, &item)
	if err != nil {
		fmt.Println("Error code : ", err)
		return
	}
	// Format totalVolume
	totalVolume := item.TotalVolume
	totalVolume = formatTotalVolume(totalVolume)

	// Mendapatkan nilai floorPrice dari API
	floorPrice := item.FloorPrice

	// Format floorPrice sesuai dengan kebutuhan
	formattedFloorPrice := formatFloorPrice(floorPrice)

	//hasil
	fmt.Println("Floor Price : ", formattedFloorPrice)
	fmt.Println("Total Volume : ", totalVolume)
	fmt.Println("Total Listed : ", item.TotalListed)
	fmt.Println("Owners : ", item.Owners)
	fmt.Println("Supply : ", item.Supply)
	fmt.Println("Pending Transactions : ", item.PendingTransactions)

}