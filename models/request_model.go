package models

type BaseModel struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
}

type Post struct {
	BaseModel
	Body string `json:"body"`
}

type Todo struct {
	BaseModel
	Completed bool `json:"completed"`
}
