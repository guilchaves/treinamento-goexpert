// INFO: 21-Instalando pacotes
package main

// go get <nome-pacote> : instala o pacote no go.mod
// go mod tidy : instala/remove pacotes na sua aplicacao que estejam sendo usados no seu codigo
import (
	"fmt"

	"github.com/google/uuid"
)

func main() {

	fmt.Println(uuid.New())
}
