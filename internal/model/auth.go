package model

type AuthInfo struct {
	Username string
	Password string
}

type userName string

const UserNameKey userName = "username"
