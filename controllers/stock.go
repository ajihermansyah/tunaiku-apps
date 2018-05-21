package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	bc "tunaiku/controllers/base"
	"tunaiku/models"
	service "tunaiku/services/stockService"

	"log"
)

type StockController struct {
	bc.BaseController
}

func (this *StockController) Index() {
	beego.SetStaticPath("/uploads", "tmp/uploads")
	this.Layout = "layouts/default.html"
	this.TplName = "stock/index.html"
	stock, err := service.GetAllStockData(&this.Service)

	if err != nil {
		log.Println("[Error] StockController.Index : ", err)
	}
	this.Data["Stock"] = stock
}

func (this *StockController) New() {
	this.TplName = "stock/form-input.html"
	this.Layout = "layouts/default.html"
	this.Data["isValid"] = true
}

func (this *StockController) Create() {
	stock := models.Stock{}
	errorMap := []string{}

	if err := this.ParseForm(&stock); err != nil {
		log.Println("[Error] StockController.Create ParseForm : ", err)

		this.Data["HasErrors"] = true
		this.Data["Errors"] = append(errorMap, "Invalid Date! Use mm/dd/yyyy format")
		this.Data["Stock"] = stock

		this.Layout = "layouts/default.html"
		this.TplName = "stock/form-input.html"
		return
	}

	valid := validation.Validation{}

	valid.Required(stock.Date, "Date").Message("is required, can't be empty")
	valid.Required(stock.Low, "Low").Message("is required, can't be empty")
	valid.Required(stock.High, "High").Message("is required, can't be empty")
	valid.Required(stock.Open, "Open").Message("is required, can't be empty")
	valid.Required(stock.Close, "Close").Message("is required, can't be empty")
	this.Data["HasErrors"] = false

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			errorMap = append(errorMap, err.Key+" "+err.Message)
		}
		this.Data["HasErrors"] = true
		this.Data["Errors"] = errorMap
		this.Data["Stock"] = stock

		this.Layout = "layouts/default.html"
		this.TplName = "stock/form-input.html"
		return
	}

	_, err := service.AddStock(&this.Service, stock)
	if err != nil {
		log.Println("[Error] StockController.Create : ", err)
		this.Redirect("/stock", 302)
		return
	}

	this.Redirect("/stock", 302)

}

func (this *StockController) CalculateStock() {
	this.Layout = "layouts/default.html"
	this.TplName = "stock/index.html"
	stockData, err := service.GetAllStockData(&this.Service)

	if err != nil {
		log.Println("[Error] StockController.Index : ", err)
	}
	this.Data["Stock"] = stockData
	stock, bestDateBuy, bestDateSell, err := service.CalculateStock(&this.Service)
	if err != nil {
		log.Println("[Error] StockController.CalculateStock : ", err)
	}
	this.Data["StockCalculate"] = stock
	this.Data["BestDateBuy"] = bestDateBuy
	this.Data["BestDateSell"] = bestDateSell

}
