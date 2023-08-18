package matriz

type matriz struct {
	id string
}

func New(id string) *matriz {
	return &matriz{id: id}
}
