package domain

import "time"

type Review struct {
	ID        int
	Version   int
	BookID    int
	UserID    int
	Rating    int
	Title     *string
	Body      *string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserName  string
}
