// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenResponse struct {
	Token string `json:"token"`
}

type UserInput struct {
	FirstName string `json:"firstName"`
	UserName  string `json:"userName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserPayload struct {
	User *User `json:"user"`
}

type UsersPayload struct {
	Users []*User `json:"users"`
	Total int     `json:"total"`
}