package mApp

import (
	"fmt"

	"github.com/8ea7b571/MoliCTF/config"
	"github.com/8ea7b571/MoliCTF/internal/mModel"
	"github.com/8ea7b571/MoliCTF/utils"
	"github.com/gin-gonic/gin"
)

type MApp struct {
	Host string
	Port uint

	secret   string
	engine   *gin.Engine
	database *mModel.MDB
}

func (mapp *MApp) Run() error {
	mapp.loadRouter()

	addr := fmt.Sprintf("%s:%d", mapp.Host, mapp.Port)
	return mapp.engine.Run(addr)
}

func NewMApp() *MApp {
	randomStr := utils.GenerateRandomString(32)

	mapp := new(MApp)
	mapp.Host = config.MConfig.MApp.Host
	mapp.Port = config.MConfig.MApp.Port
	mapp.secret = utils.MD5(randomStr)
	mapp.engine = gin.Default()
	mapp.database = mModel.NewMDB()

	return mapp
}