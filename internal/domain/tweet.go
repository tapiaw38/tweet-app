package domain

import "time"

type (
	Tweet struct {
		ID        uint64
		UserID    uint64
		User      *User
		Content   string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
