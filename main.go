package main

import (
	_ "duuit-backend/routers"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	cors "github.com/beego/beego/v2/server/web/filter/cors"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		logs.GetLogger("ORM").Println("dev mode")
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.EnableDocs = true
		beego.BConfig.WebConfig.AutoRender = true
		beego.BConfig.WebConfig.StaticDir["/v2/api/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
	}))
	beego.Run()
}