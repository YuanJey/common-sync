package main

import (
	"common-sync/internal/api/admin"
	"common-sync/internal/api/sso"
	"common-sync/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("./static/*")
	adminGroup := r.Group("/admin")
	{
		adminGroup.GET("/save_config", admin.SaveConfig)
		adminGroup.GET("/dept", admin.Dept)
		adminGroup.GET("/user", admin.User)
		//adminGroup.Use(admin.Auth)
		adminGroup.POST("/dept_map", admin.DeptMap)
		adminGroup.POST("/user_map", admin.UserMap)

	}
	ssoGroup := r.Group("/sso")
	{
		ssoGroup.GET("/oauth", sso.Oauth)
		ssoGroup.GET("/code", sso.Code)
		ssoGroup.GET("/token", sso.Token)
		ssoGroup.GET("/user_info", sso.UserInfo)
	}
	var address string
	if config.ServerConfig.Api.PORT != "" {
		address = "0.0.0.0:" + config.ServerConfig.Api.PORT
	} else {
		address = "0.0.0.0:8000"
	}
	err := r.Run(address)
	if err != nil {
		panic("api start failed " + err.Error())
	}

}
