package domain

import "time"

type Author struct {
	ID        int
	Name      string
	Bio       *string
	BirthYear *int
	CreatedAt time.Time
}

func NewAuthor(id int, name string, bio *string, birthYear *int, createdAt time.Time) Author {
	return Author{
		ID:        id,
		Name:      name,
		Bio:       bio,
		BirthYear: birthYear,
		CreatedAt: createdAt,
	}
}
