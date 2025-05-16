package dto

import "github.com/R-koma/translation-app/backend/models"

type UserResponseDto struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

// ToUserResponseDto converts a models.User instance to a UserResponseDto containing the user's ID and email.
func ToUserResponseDto(user models.User) UserResponseDto {
	return UserResponseDto{
		ID:    user.ID,
		Email: user.Email,
	}
}

// ToUserResponseDtoList converts a slice of models.User instances to a slice of UserResponseDto.
func ToUserResponseDtoList(users []models.User) []UserResponseDto {
	dtos := make([]UserResponseDto, len(users))
	for i, user := range users {
		dtos[i] = ToUserResponseDto(user)
	}
	return dtos
}
