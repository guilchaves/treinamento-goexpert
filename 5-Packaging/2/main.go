//INFO: 5.4-Entendendo go mod
package main

import "github.com/google/uuid"

func main() {
	println(uuid.New().String())
}
