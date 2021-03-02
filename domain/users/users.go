package users

type User struct {
	Id int64 `json:"id"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
	DateCreated string `json:"date_created"`
}