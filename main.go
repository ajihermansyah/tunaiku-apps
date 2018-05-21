package main

import (
	"github.com/astaxie/beego"
	"os"
	_ "tunaiku/routers"
	"tunaiku/utilities/helper"
	"tunaiku/utilities/mongo"
)

func main() {

	err := mongo.Startup(helper.MainGoRoutine)
	if err != nil {
		os.Exit(1)
	}

	beego.Run()
}
