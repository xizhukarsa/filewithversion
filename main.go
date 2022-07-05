package main

import (
	"filewithversion/api"
	"filewithversion/config"
	"filewithversion/model"
	"filewithversion/utils"
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	var cnfFile string
	flag.StringVar(&cnfFile, "f", "", "配置文件地址")
	flag.Parse()

	cnf, err := config.LoadConfig(cnfFile)
	if nil != err {
		panic(err)
	}

	db, err := cnf.ConnectDB()
	if nil != err {
		panic(err)
	}
	if err := model.AutoMigration(db); nil != err {
		panic(err)
	}

	gin.Default().Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("uuid", utils.UUID())
	})
	api.InitAPI(gin.Default())

	log.Printf("start server %v ...", cnf.Server.Port)
	if err := gin.Default().Run(fmt.Sprintf(":%v", cnf.Server.Port)); nil != err {
		panic(err)
	}
}
