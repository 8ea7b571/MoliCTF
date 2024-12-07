package mApp

import (
	"fmt"
	"net/http"

	"github.com/8ea7b571/MoliCTF/config"
	"github.com/8ea7b571/MoliCTF/internal/mModel"
	"github.com/gin-gonic/gin"
)

/*
	simple handler
*/

func (mapp *MApp) PageIndex(ctx *gin.Context) {
	resData := gin.H{
		"app": gin.H{
			"name": APP_NAME,
			"desc": APP_DESC,
			"copy": APP_COPY,
		},
	}
	ctx.HTML(http.StatusOK, "index.html", resData)
}

func (mapp *MApp) PageLogin(ctx *gin.Context) {
	resData := gin.H{
		"app": gin.H{
			"name": APP_NAME,
			"desc": APP_DESC,
			"copy": APP_COPY,
		},
	}
	ctx.HTML(http.StatusOK, "login.html", resData)
}

/*
	api handler
*/

/* user api */

func (mapp *MApp) UserLogin(ctx *gin.Context) {
	var reqData mModel.User
	err := ctx.ShouldBindJSON(&reqData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	fmt.Printf("%+v\n", reqData)
	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}

/* admin api */

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
