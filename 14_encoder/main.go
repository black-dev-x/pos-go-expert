package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number int `json:"number"`
	Amount int `json:"amount"`
}

func main() {
	account := Account{Number: 1, Amount: 1000}
	res, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}
	error := json.NewEncoder(os.Stdout).Encode(account)
	if error != nil {
		panic(error)
	}

	var accountX Account
	json.Unmarshal(res, &accountX)
	println(accountX.Amount)
}
