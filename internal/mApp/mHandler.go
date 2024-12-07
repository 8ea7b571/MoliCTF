package mApp

import (
	"net/http"

	"github.com/8ea7b571/MoliCTF/config"
	"github.com/8ea7b571/MoliCTF/internal/mModel"
	"github.com/gin-gonic/gin"
)

/*
	simple handler
*/

func (mapp *MApp) PageIndex(ctx *gin.Context) {
	userStatus := ctx.GetInt("userStatus")

	resData := gin.H{
		"app": gin.H{
			"name": APP_NAME,
			"desc": APP_DESC,
			"copy": APP_COPY,
		},
		"user": gin.H{
			"status": userStatus,
		},
	}
	ctx.HTML(http.StatusOK, "index.html", resData)
}

func (mapp *MApp) PageLogin(ctx *gin.Context) {
	userStatus := ctx.GetInt("userStatus")

	resData := gin.H{
		"app": gin.H{
			"name": APP_NAME,
			"desc": APP_DESC,
			"copy": APP_COPY,
		},
		"user": gin.H{
			"status": userStatus,
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Invalid request data.",
		})
		return
	}

	user, err := mapp.database.GetUserWithUsername(reqData.Username)
	if err != nil || user == nil || user.Password != reqData.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "Wrong username or password.",
		})
		return
	}

	jwtUser := &JwtUser{
		ID:       user.Model.ID,
		Name:     user.Name,
		Gender:   user.Gender,
		Phone:    user.Phone,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Birthday: user.Birthday,
		Username: user.Username,
		Active:   user.Active,
	}

	tokenStr, err := mapp.GenerateJwt(jwtUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Server error.",
		})
	}

	expire := config.MConfig.MApp.Expire
	ctx.SetCookie("token", tokenStr, 60*60*expire, "/", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Login success.",
	})
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
		ID:       admin.Model.ID,
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
