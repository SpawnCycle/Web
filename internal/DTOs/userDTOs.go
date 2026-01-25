package dtos

import "smaash-web/internal/models"

type UserReadDTO struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserCreateDTO struct {
	Username string `json:"username" binding:"required,max=20"`
	Email    string `json:"email" binding:"required,max=30,email"`
	Password string `json:"password" binding:"required,min=8,max=50"`
}

type UserUpdateDTO struct {
	ID       uint   `json:"id" binding:"required"`
	Username string `json:"username" binding:"required,max=20"`
	Email    string `json:"email" binding:"required,max=30,email"`
	Password string `json:"password" binding:"required,min=8,max=50"`
}

type UserLoginDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func UserToDTO(user *models.User) UserReadDTO {
	return UserReadDTO{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}
