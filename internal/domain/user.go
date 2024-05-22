package domain

type (
	User struct {
		ID        string
		FirstName string
		LastName  string
		Username  string
		Email     string
		Password  string
		CreatedAt string
		UpdatedAt string
		Followers []Follow
	}
)
