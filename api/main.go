package main

import (
	_ "api/routers"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"

	_ "github.com/lib/pq"
)

func init() {
	var dbConfig = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		"postgres",
		"12345",
		"pg",
		"teklifsts",
	)

	var maxIdle = 30 // set maximum idle connections
	var maxConn = 30 // set maximum connections (go >= 1.2)

	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", dbConfig, maxIdle, maxConn)

	// CORS Settings
	corsHandler := cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		//AllowOrigins:     []string{"http://127.0.0.1:8080", "http://*.teklifsts.tk"},
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})
	beego.InsertFilter("*", beego.BeforeRouter, corsHandler)
}

func main() {
	createTables()

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func createTables() {
	name := "default" // Database alias.
	force := false    // Drop table and re-create.
	verbose := true   // Print log.

	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}
