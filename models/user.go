package models

type User struct {
	Model
	Username	string `gorm:"column:username"		json:"username"`
	FirstName	string `gorm:"column:first_name" 	json:"first_name"`
	LastName	string `gorm:"column:last_name"		json:"last_name"`
	Email		string `gorm:"column:email"		 	json:"email"`
}

type CreateUserInput struct {
	Username 	string `json:"username"		binding:"required"`
	FirstName 	string `json:"first_name"	binding:"required"`
	LastName 	string `json:"last_name"	binding:"required"`
	Email 		string `json:"email" 		binding:"requried"`
}

type UpdateUserInput struct {
	Username 	string `json:"username"`
	FirstName 	string `json:"first_name"`
	LastName 	string `json:"last_name"`
	Email 		string `json:"email"`
}