package models

type AccessToken struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
}
