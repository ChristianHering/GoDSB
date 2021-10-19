package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type BitcoinPrice struct {
	Data struct {
		Amount   string `json:"amount"`
		Currency string `json:"currency"`
	} `json:"data"`
}

func bitcoinEventCallback(position *int, button *int) (f callbackFunc, err error) {
	return nil, nil
}

//Don't call this function more than once per second
//or you may end up getting rate limited by coinbase
func bitcoinCallback(position int) error {
	//Get the price of btc from 7 days ago
	response, err := http.Get("https://api.coinbase.com/v2/prices/spot?currency=USD&date=" + time.Now().Add(-7*24*time.Hour).Format("2006-01-02"))
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var dayPrice = BitcoinPrice{}

	err = json.Unmarshal(body, &dayPrice)
	if err != nil {
		return err
	}

	dayPriceValue, err := strconv.ParseFloat(dayPrice.Data.Amount, 64)
	if err != nil {
		return err
	}

	//Get the current price of btc
	response, err = http.Get("https://api.coinbase.com/v2/prices/spot?currency=USD")
	if err != nil {
		return err
	}

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var currentPrice = BitcoinPrice{}

	err = json.Unmarshal(body, &currentPrice)
	if err != nil {
		return err
	}

	currentPriceValue, err := strconv.ParseFloat(currentPrice.Data.Amount, 64)
	if err != nil {
		return err
	}

	//Calculate the % change between the two prices
	change := strconv.FormatFloat((((currentPriceValue - dayPriceValue) / dayPriceValue) * 100), 'f', 2, 64)

	if dayPriceValue > currentPriceValue {
		sectionStatus[position] = "BTC/USD: $" + currentPrice.Data.Amount + " " + change + "% 🔻"
	} else {
		sectionStatus[position] = "BTC/USD: $" + currentPrice.Data.Amount + " +" + change + "% 🚀"
	}

	return nil
}
