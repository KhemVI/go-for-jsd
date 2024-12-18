package common

type CreateUserRequest struct {
	Name        string `binding:"required" example:"Khem"`
	Email       string `binding:"required" example:"dog@gmail.com"`
	Password    string `binding:"required" example:"1234567"`
	MyCustomKey int    `binding:"required" example:"1234"`
}
