package Resrik

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListItem struct {
	Id     int
	ListId int
	ItemId int
}
