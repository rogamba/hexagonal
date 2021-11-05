package models

type Tweet struct {
	UserId string `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
