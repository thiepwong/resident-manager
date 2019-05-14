package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kataras/iris"
	"github.com/thiepwong/resident-manager/common"
	"github.com/thiepwong/resident-manager/routes"
)

func main() {
	_cfgPath := "./configs/app.yaml"
	conf, es := common.LoadConfig(_cfgPath)
	if es != nil {
		os.Exit(10)
	}
	app := iris.Default()

	crs := func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin,Content-Type")
		ctx.Next()
	}

	app.Logger().SetLevel("debug")
	routes.RegisterRoute(app, crs, conf)

	er := app.Run(iris.Addr(conf.Service.Host+":"+strconv.Itoa(conf.Service.Port)), iris.WithoutPathCorrectionRedirection)
	if er != nil {
		fmt.Println("Server not started!")
	}

}
