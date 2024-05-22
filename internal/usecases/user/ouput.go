package user

import (
	"time"

	"github.com/tapiaw38/tweet-app/internal/domain"
)

type (
	UserOutputData struct {
		ID        uint64    `json:"id,omitempty"`
		FirstName string    `json:"first_name,omitempty"`
		LastName  string    `json:"last_name,omitempty"`
		Email     string    `json:"email,omitempty"`
		CreatedAt time.Time `json:"created_at,omitempty"`
		UpdatedAt time.Time `json:"updated_at,omitempty"`
	}
)

func toUserOutput(user domain.User) UserOutputData {
	return UserOutputData{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
