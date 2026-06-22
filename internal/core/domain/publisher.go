package domain

type Publisher struct {
	ID   int
	Name string
}

func NewPublisher(id int, name string) Publisher {
	return Publisher{
		ID:   id,
		Name: name,
	}
}
