package models

type User struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type News struct {
	ID    int    `json:"-" db:"id"`
	Title string `json:"title" db:"title"`
	Body  string `json:"body" db:"body"`
	Image string `json:"image" db:"image"`
	Mark  string `json:"mark" db:"mark"`
	Reliz string `json:"reliz" db:"reliz"`
}

type UpdateNews struct {
	Title *string `json:"title"`
	Body  *string `json:"body"`
	Image *string `json:"image"`
	Mark  *string `json:"mark"`
	Reliz *string `json:"reliz"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
