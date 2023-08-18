package matriz

type matriz struct {
	id     string
	codigo int
}

func New(id string) *matriz {
	return &matriz{id: id}
}
