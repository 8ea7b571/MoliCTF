package mApp

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/8ea7b571/MoliCTF/config"
	"github.com/8ea7b571/MoliCTF/internal/mModel"
	"github.com/gin-gonic/gin"
)

func (mapp *MApp) UserRegister(ctx *gin.Context) {
	var preUser *mModel.User
	var reqData map[string]interface{}

	err := ctx.ShouldBindJSON(&reqData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Invalid request data",
		})
		return
	}

	firstname := reqData["firstname"]
	lastname := reqData["lastname"]
	if firstname == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Firstname is required",
		})
		return
	}

	_gender := reqData["gender"]
	if _gender == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Gender is required",
		})
		return
	}
	gender, _ := strconv.ParseUint(_gender.(string), 10, 64)

	phone := reqData["phone"]
	if phone == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Phone is required",
		})
		return
	}

	// check if the phone is already used
	preUser, _ = mapp.database.GetUserWithPhone(phone.(string))
	if preUser != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Phone already used",
		})
		return
	}

	email := reqData["email"]
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Email is required",
		})
		return
	}

	// check if the email is already used
	preUser, _ = mapp.database.GetUserWithEmail(email.(string))
	if preUser != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Email already used",
		})
		return
	}

	username := reqData["username"]
	if username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Username is required",
		})
		return
	}

	// check if the username is already used
	preUser, _ = mapp.database.GetUserWithUsername(username.(string))
	if preUser != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Username already used",
		})
		return
	}

	password1 := reqData["password1"]
	if password1 == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Password is required",
		})
		return
	}

	password2 := reqData["password2"]
	if password2 == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Password confirm is required",
		})
		return
	}

	if password1 != password2 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Passwords do not match",
		})
		return
	}

	user := &mModel.User{
		Name:         strings.TrimSpace(firstname.(string) + " " + lastname.(string)),
		Gender:       uint(gender),
		Phone:        phone.(string),
		Email:        email.(string),
		Avatar:       "/upload/images/default-avatar.jpg",
		Introduction: "No introduction yet.",
		Username:     username.(string),
		Password:     password1.(string),
		Active:       true,
		Score:        0,
		TeamId:       0,
	}

	_, err = mapp.database.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Register success",
	})
}

func (mapp *MApp) UserLogin(ctx *gin.Context) {
	var reqData mModel.User
	err := ctx.ShouldBindJSON(&reqData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Invalid request data",
		})
		return
	}

	user, err := mapp.database.GetUserWithUsername(reqData.Username)
	if err != nil || user == nil || user.Password != reqData.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "Wrong username or password",
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
		Username: user.Username,
		Active:   user.Active,
	}

	tokenStr, err := mapp.GenerateJwt(jwtUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Server error",
		})
		return
	}

	// set user token to cache
	mapp.cache.User.Set(jwtUser.Username, tokenStr)

	expire := config.MConfig.MApp.Expire
	ctx.SetCookie("token", tokenStr, 60*60*expire, "/", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Login success",
	})
}

func (mapp *MApp) UserLogout(ctx *gin.Context) {
	jwtUser := ctx.MustGet("jwtUser").(*JwtUser)

	// remove user token from cache
	mapp.cache.User.Del(jwtUser.Username)
	ctx.Redirect(http.StatusFound, "/")
}
