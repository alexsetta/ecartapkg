package objeto

type Objeto struct {
	id string
}

func New(id string) *Objeto {
	return &Objeto{id: id}
}
