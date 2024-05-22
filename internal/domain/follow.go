package domain

type (
	Follow struct {
		ID        uint64
		UserID    uint64
		FollowID  uint64
		CreatedAt string
	}
)
