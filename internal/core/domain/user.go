package domain

import "time"

type User struct {
	ID      int
	Version int

	Email        string
	PasswordHash string
	FullName     string
	PhoneNumber  *string
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewUser(
	id int,
	version int,
	email string,
	passwordHash string,
	fullName string,
	phoneNumber *string,
	role string,
	createdAt time.Time,
	updatedAt time.Time,
) User {
	return User{
		ID:           id,
		Version:      version,
		Email:        email,
		PasswordHash: passwordHash,
		FullName:     fullName,
		PhoneNumber:  phoneNumber,
		Role:         role,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}

func NewUserUninitialized(email string, passwordHash string, fullName string, phoneNumber *string) User {
	return NewUser(UninitializedID, UninitializedVersion, email, passwordHash, fullName, phoneNumber, "customer", time.Now(), time.Now())
}
