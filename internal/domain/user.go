package domain

import "time"

type (
	User struct {
		ID        uint64
		FirstName string
		LastName  string
		Username  string
		Email     string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
