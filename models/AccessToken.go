package models

type AccessToken struct {
	AccessToken    string
	RefreshToken   string
	ExpirationTime int64
}
