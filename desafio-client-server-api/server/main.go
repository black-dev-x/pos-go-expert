package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"black.dev.x/database"
	"gorm.io/gorm"
)

type DolarPrice struct {
	gorm.Model
	Code               string `json:"code"`
	Codein             string `json:"codein"`
	Name               string `json:"name"`
	High               string `json:"high"`
	Low                string `json:"low"`
	VarBid             string `json:"varBid"`
	PctChange          string `json:"pctChange"`
	Bid                string `json:"bid"`
	Ask                string `json:"ask"`
	ExternalTimestamp  string `json:"timestamp"`
	ExternalCreateDate string `json:"create_date"`
}

func main() {
	migrate()
	http.HandleFunc("/", getDolarPrice)
	http.ListenAndServe(":8080", nil)

}

func getDolarPrice(w http.ResponseWriter, r *http.Request) {
	ctxApi := context.Background()
	ctxApi, cancel := context.WithTimeout(ctxApi, time.Millisecond*200)
	defer cancel()
	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	resp, err := http.NewRequestWithContext(ctxApi, "GET", url, nil)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	price := DolarPrice{}
	error := json.NewDecoder(resp.Body).Decode(&price)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(price)

}

func migrate() {
	database.Load()
	database.DB.AutoMigrate(&DolarPrice{})
}
