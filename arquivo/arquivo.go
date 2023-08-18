package arquivo

type arquivo struct {
	id   string
	nome string
}

func New(id string) *arquivo {
	return &arquivo{id: id}
}
