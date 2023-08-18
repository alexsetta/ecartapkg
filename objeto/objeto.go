package objeto

type Objeto struct {
	id     string
	codigo int
}

func New(id string) *Objeto {
	return &Objeto{id: id}
}
