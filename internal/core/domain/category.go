package domain

type Category struct {
	ID          int
	Name        string
	Slug        string
	Description *string
}

func NewCategory(id int, name, slug string, description *string) Category {
	return Category{
		ID:          id,
		Name:        name,
		Slug:        slug,
		Description: description,
	}
}
