package cache

import "acai-go/models"

var oauthCache map[string]models.User

var publicKey = ""

func GetPublicKey() string {
	return publicKey
}

func SetPublicKey(key string) {
	publicKey = key
}

func GetAccessToken(key string) models.User {
	return oauthCache[key]
}

func SetAccessToken(key string, user models.User) {
	oauthCache[key] = user
}
