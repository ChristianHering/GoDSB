package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type EtherMineCurrentStats struct {
	Status string `json:"status"`
	Data   struct {
		Time             int         `json:"time"`
		Lastseen         int         `json:"lastSeen"`
		Reportedhashrate float64     `json:"reportedHashrate"`
		Currenthashrate  float64     `json:"currentHashrate"`
		Validshares      int         `json:"validShares"`
		Invalidshares    int         `json:"invalidShares"`
		Staleshares      int         `json:"staleShares"`
		Averagehashrate  float64     `json:"averageHashrate"`
		Activeworkers    int         `json:"activeWorkers"`
		Unpaid           int64       `json:"unpaid"`
		Unconfirmed      interface{} `json:"unconfirmed"`
		Coinspermin      float64     `json:"coinsPerMin"`
		Usdpermin        float64     `json:"usdPerMin"`
		Btcpermin        float64     `json:"btcPerMin"`
	} `json:"data"`
}

func miningEventCallback(position *int, button *int) (f callbackFunc, err error) {

	return nil, nil
}

//Information is updated every other minute but
//the rate limit's 100 requests/15 minutes/IP
func miningCallback(position int) error {
	response, err := http.Get("https://api.ethermine.org/miner/:0d1Ff17e61A58066508aB6a78dc4E73503B02E6C/currentStats") //I accept donations ;D
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var ethermineData = EtherMineCurrentStats{}

	err = json.Unmarshal(body, &ethermineData)
	if err != nil {
		return err
	}

	averageProfit := ethermineData.Data.Usdpermin * 60.0 * 24.0 * 30.0

	profit := (averageProfit * ethermineData.Data.Reportedhashrate) / ethermineData.Data.Averagehashrate

	//sectionStatus[position] = "$" + strconv.FormatFloat(averageProfit, 'f', 2, 64) + " avg  $" + strconv.FormatFloat(profit, 'f', 2, 64) + " inst/mo 💰"
	sectionStatus[position] = "$" + strconv.FormatFloat(profit, 'f', 2, 64) + "/mo 💰"

	return nil
}
