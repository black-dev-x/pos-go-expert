package main

import (
	"encoding/json"
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

	http.HandleFunc("/", FindAddress)
	http.ListenAndServe(":8080", nil)
}

func FindAddress(w http.ResponseWriter, r *http.Request) {
	zipCode := r.URL.Query().Get("cep")
	if zipCode == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("CEP is required"))
		return
	}
	req, err := http.Get("https://viacep.com.br/ws/" + zipCode + "/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()
	address := Address{}
	err = json.NewDecoder(req.Body).Decode(&address)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(address)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(address)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)

}
