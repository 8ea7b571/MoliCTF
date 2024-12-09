package mApp

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

func (mapp *MApp) PageRegister(ctx *gin.Context) {
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
	ctx.HTML(http.StatusOK, "register.html", resData)
}

func (mapp *MApp) PageUsers(ctx *gin.Context) {
	type resUser struct {
		Id     uint   `json:"id"`
		Name   string `json:"name"`
		Gender string `json:"gender"`
		Phone  string `json:"phone"`
		Email  string `json:"email"`
		Team   string `json:"team"`
		Score  uint   `json:"score"`
	}

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit := 15
	offset := (page - 1) * limit

	userStatus := ctx.GetInt("userStatus")
	userPtrList, err := mapp.database.GetUsers(offset, limit)
	userCount, err := mapp.database.GetUserCount()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Server error",
		})
		return
	}

	userList := make([]resUser, len(userPtrList))
	for i, userPtr := range userPtrList {
		if userPtr != nil {
			teamPtr, _ := mapp.database.GetTeamWithId(userPtr.TeamId)
			userList[i] = resUser{
				Id:    userPtr.ID,
				Name:  userPtr.Name,
				Phone: userPtr.Phone,
				Email: userPtr.Email,
				Score: userPtr.Score,
				Team:  teamPtr.Name,
				Gender: func(gender uint) string {
					if userPtr.Gender == 1 {
						return "Male"
					} else {
						return "Female"
					}
				}(userPtr.Gender),
			}
		}
	}

	fmt.Println((userCount + limit - 1) / limit)

	resData := gin.H{
		"app": gin.H{
			"name": APP_NAME,
			"desc": APP_DESC,
			"copy": APP_COPY,
		},
		"user": gin.H{
			"status": userStatus,
			"list":   userList,
		},
		"data": gin.H{
			"currentPage": page,
			"totalPage":   (userCount + limit - 1) / limit,
		},
	}
	ctx.HTML(http.StatusOK, "users.html", resData)
}

/*
	api handler
*/

/* user api */

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
		Name:     strings.TrimSpace(firstname.(string) + " " + lastname.(string)),
		Gender:   uint(gender),
		Phone:    phone.(string),
		Email:    email.(string),
		Avatar:   "/upload/images/default-avatar.jpg",
		Username: username.(string),
		Password: password1.(string),
		Active:   true,
		Score:    0,
		TeamId:   0,
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
