package objeto

type Objeto struct {
	id      string
	version string
}

func New(id string) *Objeto {
	return &Objeto{id: id, version: "1.4.2"}
}
