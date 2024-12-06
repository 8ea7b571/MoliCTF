package mApp

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var skipPaths = []string{
	"/v1/admin/login",
}

func (mapp *MApp) loadMiddleware() {
	mapp.engine.Use(mapp.jwtAuthMiddleware())
}

func (mapp *MApp) jwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, skipPath := range skipPaths {
			if skipPath == ctx.FullPath() {
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
