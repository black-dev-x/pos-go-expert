package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Address struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	zipCode := "74543490"
	req, err := http.Get("https://viacep.com.br/ws/" + zipCode + "/json")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	address := Address{}
	err = json.NewDecoder(req.Body).Decode(&address)
	if err != nil {
		panic(err)
	}
	fmt.Println(address)
}
