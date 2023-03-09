package models

type User struct {
	Id      string `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	ADDRESS string `json:"address"`
	Age     int    `json:"age"`
}

type CreateUser struct {
	Name    string `json:"name" binding:"required"`
	ADDRESS string `json:"address" binding:"required"`
	Age     int    `json:"age" binding:"required"`
}

type UpdateUser struct {
	Name    string `json:"name"`
	ADDRESS string `json:"address"`
	Age     int    `json:"age"`
}
