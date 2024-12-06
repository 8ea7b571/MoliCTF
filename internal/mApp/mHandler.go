package mApp

import (
	"net/http"

	"github.com/8ea7b571/MoliCTF/config"
	"github.com/8ea7b571/MoliCTF/internal/mModel"
	"github.com/gin-gonic/gin"
)

func (mapp *MApp) AdminLogin(ctx *gin.Context) {
	var reqData mModel.Admin
	err := ctx.ShouldBindJSON(&reqData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	admin, err := mapp.database.GetAdminWithUsername(reqData.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	if admin == nil || reqData.Password != admin.Password {
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}

	adminJwt, err := GenerateJwtForAdmin(admin, mapp.secret)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}

	expire := config.MConfig.MApp.Expire
	
	ctx.SetCookie("token", adminJwt, 60*60*expire, "/", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"admin": admin})
}
