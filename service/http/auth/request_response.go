package auth

import "github.com/Mrhb787/hospital-ward-manager/model"

type SignInUserRequest struct {
	PhoneNumber string
	Password    string
}

type SignInUserResponse struct {
	User  model.User
	Token string
}

type ValidateTokenRequest struct {
	Token string
}

type ValidateTokenResponse struct {
	IsValid bool
	Session model.UserSession
}
