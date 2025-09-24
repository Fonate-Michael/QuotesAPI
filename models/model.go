package models

type Quote struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

type Comment struct {
	ID      int    `json:"id"`
	User_id int    `json:"user_id"`
	Comment string `json:"comment"`
}
