package arquivo

type arquivo struct {
	id string
}

func New(id string) *arquivo {
	return &arquivo{id: id}
}
