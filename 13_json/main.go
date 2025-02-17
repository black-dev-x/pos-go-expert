package main

import "encoding/json"

type Account struct {
	Number int
	Amount int
}

func main() {
	account := Account{Number: 1, Amount: 1000}
	res, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}
	println(string(res))
}
