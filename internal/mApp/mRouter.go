package mApp

func (mapp *MApp) loadRouter() {
	v1 := mapp.engine.Group("v1")
	{
		admin := v1.Group("admin")
		{
			admin.POST("/login", mapp.AdminLogin)
		}
	}
}
