package mApp

import (
	"fmt"

	"github.com/8ea7b571/MoliCTF/config"
	"github.com/8ea7b571/MoliCTF/internal/mCache"
	"github.com/8ea7b571/MoliCTF/internal/mModel"
	"github.com/8ea7b571/MoliCTF/utils"
	"github.com/gin-gonic/gin"
)

type MApp struct {
	Host string
	Port uint

	secret   string
	cache    *mCache.MCache
	engine   *gin.Engine
	database *mModel.MDB
}

func (mapp *MApp) loadMiddleware() {
	mapp.engine.Use(mapp.jwtAuthMiddleware())
}

func (mapp *MApp) loadTemplates() {
	assetsPath := fmt.Sprintf("%s/assets", config.MConfig.MApp.Template)
	uploadPath := fmt.Sprintf("%s/data/upload", config.MConfig.MApp.Root)
	htmlPath := fmt.Sprintf("%s/html/*.html", config.MConfig.MApp.Template)

	mapp.engine.Static("/assets", assetsPath)
	mapp.engine.Static("/upload", uploadPath)
	mapp.engine.LoadHTMLGlob(htmlPath)
}

func (mapp *MApp) Run() error {
	mapp.loadMiddleware()
	mapp.loadRouter()
	mapp.loadTemplates()

	addr := fmt.Sprintf("%s:%d", mapp.Host, mapp.Port)
	return mapp.engine.Run(addr)
}

func NewMApp() *MApp {
	randomStr := utils.GenerateRandomString(32)

	mapp := new(MApp)
	mapp.Host = config.MConfig.MApp.Host
	mapp.Port = config.MConfig.MApp.Port
	mapp.secret = utils.MD5(randomStr)
	mapp.cache = mCache.NewMCache()
	mapp.engine = gin.Default()
	mapp.database = mModel.NewMDB()

	return mapp
}
