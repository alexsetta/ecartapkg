package matriz

import "fmt"

type matriz struct {
	codigo int
}

func New() *matriz {
	fmt.Println("Criando uma nova matriz")
	return &matriz{}
}
