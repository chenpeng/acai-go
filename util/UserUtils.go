package util

import (
	"acai-go/cache"
	"acai-go/models"
	"github.com/astaxie/beego/context"
)

func GetUser(ctx *context.Context) models.User {
	accessToken := ctx.Input.Header("accessToken")
	user := cache.GetAccessToken(accessToken)
	return user
}
