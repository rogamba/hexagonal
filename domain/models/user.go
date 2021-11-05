package models

type User struct {
	UserId   string  `json:"user_id"`
	Name     string  `json:"name"`
	LastName string  `json:"last_name"`
	Age      int64   `json:"age"`
	Address  string  `json:"address"`
	Tweets   []Tweet `json:"tweets"`
}
