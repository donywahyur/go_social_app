package model

type Follower struct {
	UserID     string `json:"user_id"`
	FollowerID string `json:"follower_id"`
	CreatedAt  string `json:"created_at"`
	User       User   `json:"user"`
	Follower   User   `json:"follower"`
}
