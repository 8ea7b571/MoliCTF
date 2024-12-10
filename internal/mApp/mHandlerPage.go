package mApp

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
		Id           uint   `json:"id"`
		Name         string `json:"name"`
		Gender       string `json:"gender"`
		Introduction string `json:"introduction"`
		Email        string `json:"email"`
		Team         string `json:"team"`
		Score        uint   `json:"score"`
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
				Id:           userPtr.ID,
				Name:         userPtr.Name,
				Score:        userPtr.Score,
				Team:         teamPtr.Name,
				Email:        userPtr.Email,
				Introduction: userPtr.Introduction,
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

	resData := gin.H{
		"app": gin.H{
			"name": APP_NAME,
			"desc": APP_DESC,
			"copy": APP_COPY,
		},
		"user": gin.H{
			"status": userStatus,
		},
		"data": gin.H{
			"currentPage": page,
			"totalPage":   (userCount + limit - 1) / limit,
			"userList":    userList,
		},
	}
	ctx.HTML(http.StatusOK, "users.html", resData)
}
