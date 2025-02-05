// INFO: 5-Busca cep
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
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
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while making the request: %v\n", err)
		}
		defer resp.Body.Close()
		content, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while reading the response: %v\n", err)
		}

		var data ViaCEP
		err = json.Unmarshal(content, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while parsing the response: %v\n", err)
		}

		fmt.Println(data.Localidade)

		file, err := os.Create("city.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while creating the file: %v\n", err)
		}
		defer file.Close()

		_, err = file.WriteString(
			fmt.Sprintf("CEP: %s, Logradouro: %s, Localidade: %s, UF: %s\n", data.Cep, data.Logradouro, data.Localidade, data.Uf),
		)

		fmt.Println("File created successfully")
		fmt.Println("CEP: %s, Logradouro: %s, Localidade: %s, UF: %s\n", data.Cep, data.Logradouro, data.Localidade, data.Uf)
	}

}
