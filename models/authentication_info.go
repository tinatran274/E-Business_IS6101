package models

var AuthInfoKey = "auth_info"

type AuthenticationInfo struct {
	User *User `json:"user"`
}

func NewAuthenticator(user *User) *AuthenticationInfo {
	return &AuthenticationInfo{
		User: user,
	}
}

