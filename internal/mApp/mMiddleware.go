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
	"/v1/user/login",
	"/v1/admin/login",
}

func (mapp *MApp) jwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if strings.HasPrefix(ctx.Request.URL.Path, "/assets") {
			ctx.Next()
			return
		}

		for _, skipPath := range skipPaths {
			if skipPath == ctx.Request.URL.Path {
				ctx.Next()
				return
			}
		}

		tokenRaw := strings.Split(ctx.GetHeader("Authorization"), " ")
		if len(tokenRaw) != 2 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenStr := tokenRaw[1]
		if tokenStr == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		jwtUser, err := mapp.ParseJwt(tokenStr)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("jwtUser", jwtUser)
		ctx.Next()
	}
}
