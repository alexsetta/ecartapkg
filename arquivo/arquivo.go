package arquivo

import "fmt"

type arquivo struct {
	nome string
}

func New() *arquivo {
	fmt.Println("Criando um novo arquivo")
	return &arquivo{}
}
