package routers

import (
	"github.com/astaxie/beego"
	"tunaiku/controllers"
)

func init() {
	beego.Router("/stock", &controllers.StockController{}, "get:Index")
	beego.Router("/stock/input-data", &controllers.StockController{}, "get:New")
	beego.Router("/stock/input-data", &controllers.StockController{}, "post:Create")
	beego.Router("/stock/calculate-stock", &controllers.StockController{}, "get:CalculateStock")
}
