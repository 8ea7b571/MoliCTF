package mApp

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var skipPaths = []string{
	"/",
	"/score",
	"/users",
	"/teams",
	"/login",
	"/register",
	"/challenges",
	"/v1/user/login",
	"/v1/user/logout",
}

func (mapp *MApp) jwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// skip static files
		if strings.HasPrefix(ctx.Request.URL.Path, "/assets") {
			ctx.Next()
			return
		}

		ctx.Set("userStatus", 0)
		cToken, err := ctx.Cookie("token")
		if err == nil || cToken != "" {
			jwtUser, err := mapp.ParseJwt(cToken)
			if err == nil && mapp.cache.User.Get(jwtUser.Username) != nil {
				ctx.Set("jwtUser", jwtUser)
				ctx.Set("userStatus", 1)
			}
		}

		// skip whitelist paths
		for _, skipPath := range skipPaths {
			if skipPath == ctx.Request.URL.Path {
				ctx.Next()
				return
			}
		}

		if ctx.GetInt("userStatus") == 1 {
			ctx.Next()
		} else {
			ctx.Redirect(http.StatusMovedPermanently, "/login")
		}
	}
}
