package filter

import (
	"acai-go/cache"
	"acai-go/dto"
	"github.com/astaxie/beego/context"
)

var urlList = []string{
	"/api/oauth/publicKey",
	"/api/oauth/signUp",
	"/api/oauth/signIn",
}

var OauthFilter = func(ctx *context.Context) {
	url := ctx.Request.RequestURI
	if !contains(urlList, url) {
		accessToken := ctx.Input.Header("accessToken")
		if accessToken == "" {
			result := dto.Result{
				Code:    -1,
				Data:    nil,
				Message: "请登录",
			}
			ctx.Output.JSON(result, true, true)
		} else {
			user := cache.GetAccessToken(accessToken)
			if user.Id == 0 {
				result := dto.Result{
					Code:    -1,
					Data:    nil,
					Message: "请登录",
				}
				ctx.Output.JSON(result, true, true)
			}
		}
	}
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
