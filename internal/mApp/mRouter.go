package mApp

func (mapp *MApp) loadRouter() {
	/* simple route */
	mapp.engine.GET("/", mapp.PageIndex)
	mapp.engine.GET("/login", mapp.PageLogin)
	mapp.engine.GET("/register", mapp.PageRegister)

	/* api route */
	v1 := mapp.engine.Group("v1")
	{
		user := v1.Group("/user")
		{
			user.POST("/login", mapp.UserLogin)
			user.POST("/register", mapp.UserRegister)
			user.GET("/logout", mapp.UserLogout)
		}
	}
}
