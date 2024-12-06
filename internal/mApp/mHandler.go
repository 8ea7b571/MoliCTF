package mApp

import (
	"net/http"

	"github.com/8ea7b571/MoliCTF/config"
	"github.com/8ea7b571/MoliCTF/internal/mModel"
	"github.com/gin-gonic/gin"
)

/* simple handler */

func (mapp *MApp) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{"page": "index"})
}

/* api handler */

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

	jwtUser := &JwtUser{
		Name:     admin.Name,
		Gender:   admin.Gender,
		Phone:    admin.Phone,
		Email:    admin.Email,
		Avatar:   admin.Avatar,
		Birthday: admin.Birthday,
		Username: admin.Username,
		Active:   admin.Active,
	}

	tokenStr, err := mapp.GenerateJwt(jwtUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}

	expire := config.MConfig.MApp.Expire
	ctx.SetCookie("token", tokenStr, 60*60*expire, "/", "", false, true)

	// TODO: redirect to admin panel
	ctx.Redirect(http.StatusFound, "/")
}

func (mapp *MApp) AdminInfo(ctx *gin.Context) {
	jwtUser := ctx.MustGet("jwtUser").(*JwtUser)
	admin, err := mapp.database.GetAdminWithUsername(jwtUser.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}

	admin.Password = ""
	ctx.JSON(http.StatusOK, gin.H{"data": admin})
}
