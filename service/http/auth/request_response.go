package auth

import "github.com/mahima-c/fruito/model"

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
