package user

import "github.com/tapiaw38/tweet-app/internal/domain"

type (
	UserInputData struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	FollowInputData struct {
		UserID   uint64 `json:"user_id"`
		FollowID uint64 `json:"follow_id"`
	}
)

func toUserInput(user UserInputData) domain.User {
	return domain.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
	}
}

func toFollowInput(follower FollowInputData) domain.Follow {
	return domain.Follow{
		UserID:   follower.UserID,
		FollowID: follower.FollowID,
	}
}
