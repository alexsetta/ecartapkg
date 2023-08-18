package objeto

import "fmt"

type Objeto struct {
	codigo int
}

func New() *Objeto {
	fmt.Println("Criando um novo objeto")
	return &Objeto{}
}
