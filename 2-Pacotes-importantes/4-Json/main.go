// INFO: 2.4-Trabalhando com Json
package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number  int `json:"account_number"`
	Balance int `json:"balance"`
}

func main() {
	account := Account{Number: 1, Balance: 1000}
	res, err := json.Marshal(account)
	if err != nil {
		println(err)
	}
	println(string(res))

	// same as above but not saving the result inside a variable
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(account)
	if err != nil {
		println(err)
	}

	var account2 Account
	rawJson := []byte(`{"account_balance":2,"balance":2000}`)
	err = json.Unmarshal(rawJson, &account2)
	if err != nil {
		println(err)
	}

	println(account2.Balance)

}
