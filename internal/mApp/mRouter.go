package mApp

func (mapp *MApp) loadRouter() {
	/* simple route */
	mapp.engine.GET("/", mapp.Index)

	/* api route */
	v1 := mapp.engine.Group("v1")
	{
		admin := v1.Group("admin")
		{
			admin.GET("/info", mapp.AdminInfo)
			admin.POST("/login", mapp.AdminLogin)
		}
	}
}
