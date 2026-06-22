package domain

import "time"

type Address struct {
	ID            int
	Version       int
	UserID        int
	StreetAddress string
	City          string
	PostalCode    string
	Country       string
	IsDefault     bool
	CreatedAt     time.Time
}

func NewAddress(
	id int,
	version int,
	userID int,
	streetAddress string,
	city string,
	postalCode string,
	country string,
	isDefault bool,
	createdAt time.Time,
) Address {
	return Address{
		ID:            id,
		Version:       version,
		UserID:        userID,
		StreetAddress: streetAddress,
		City:          city,
		PostalCode:    postalCode,
		Country:       country,
		IsDefault:     isDefault,
		CreatedAt:     createdAt,
	}
}

func NewAddressUninitialized(
	userID int,
	streetAddress string,
	city string,
	postalCode string,
	country string,
	isDefault bool,
) Address {
	return Address{
		UninitializedID,
		UninitializedVersion,
		userID,
		streetAddress,
		city,
		postalCode,
		country,
		isDefault,
		time.Now().UTC(),
	}
}
