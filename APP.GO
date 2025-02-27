package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type CryptoData struct {
	ID     string  `json:"id"`
	Symbol string  `json:"symbol"`
	Name   string  `json:"name"`
	Price  float64 `json:"current_price"`
}

const apiURL = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=brl&ids=bitcoin,ethereum,dogecoin"

func fetchCryptoData() ([]CryptoData, error) {
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var cryptos []CryptoData
	if err := json.NewDecoder(resp.Body).Decode(&cryptos); err != nil {
		return nil, err
	}
	return cryptos, nil
}

func saveJSONToFile(data []CryptoData) error {
	file, err := os.Create("cryptos.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // formata com indentação
	return encoder.Encode(data)
}

func cryptoHandler(w http.ResponseWriter, r *http.Request) {
	// Busca os dados
	cryptos, err := fetchCryptoData()
	if err != nil {
		http.Error(w, "Erro ao buscar dados: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Exibir os preços no terminal
	fmt.Println("📊 Preços das Criptomoedas:")
	for _, crypto := range cryptos {
		fmt.Printf("%s (%s): R$%.2f\n", crypto.Name, crypto.Symbol, crypto.Price)
	}

	// Salva os dados em um arquivo JSON
	if err := saveJSONToFile(cryptos); err != nil {
		log.Println("Erro ao salvar JSON:", err)
	}

	// Configura o header e retorna o JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cryptos); err != nil {
		http.Error(w, "Erro ao enviar resposta: "+err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/cryptos", cryptoHandler)

	fmt.Println("Servidor rodando na porta 8080. Acesse http://localhost:8080/cryptos")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Erro no servidor:", err)
	}
}
