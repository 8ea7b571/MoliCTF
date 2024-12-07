package mApp

func (mapp *MApp) loadRouter() {
	/* simple route */
	mapp.engine.GET("/", mapp.PageIndex)
	mapp.engine.GET("/login", mapp.PageLogin)

	/* api route */
	v1 := mapp.engine.Group("v1")
	{
		user := v1.Group("/user")
		{
			user.POST("/login", mapp.UserLogin)
		}

		admin := v1.Group("admin")
		{
			admin.POST("/login", mapp.AdminLogin)
		}
	}
}
