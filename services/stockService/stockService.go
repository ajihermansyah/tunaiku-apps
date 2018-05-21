package stockService

import (
	"gopkg.in/mgo.v2"

	"log"

	"github.com/astaxie/beego"

	"math"
	"tunaiku/models"
	"tunaiku/services"
)

func GetAllStockData(service *services.Service) (*[]models.Stock, error) {
	stock := []models.Stock{}

	f := func(collection *mgo.Collection) error {
		return collection.Find(nil).Sort("date").All(&stock)
	}

	if err := service.DBAction(beego.AppConfig.String("mgo_database"), "stock", f); err != nil {
		log.Println("[Error] StockService.GetAllStockData : ", err)
		if err != mgo.ErrNotFound {
			return nil, err
		}
	}

	for i, item := range stock {
		stock[i].DateStr = item.Date.Format("02/01/2006")
	}

	return &stock, nil
}

func AddStock(service *services.Service, stock models.Stock) (*models.Stock, error) {
	f := func(collection *mgo.Collection) error {
		return collection.Insert(&stock)
	}

	if err := service.DBAction(beego.AppConfig.String("mgo_database"), "stock", f); err != nil {
		log.Println("[Error] StockService.AddStock : ", err)
		return &stock, err
	}

	return &stock, nil
}

func CalculateStock(service *services.Service) (*[]models.Stock, string, string, error) {
	stock, err := GetAllStockData(service)
	if err != nil {
		log.Println("[Error] StockService.CalculateStock GetAllStockData", err)
		return nil, "", "", err
	}

	stocks := *stock
	buyDate := stocks[0].Date
	sellDate := stocks[1].Date
	buyValue := math.MaxInt32
	sellValue := 0

	for i, item := range stocks {
		temp1 := 0
		temp2 := 0
		temp1 = item.High - item.Close
		temp2 = item.Close - item.Low
		if temp1 < 0 {
			temp1 = temp1 * -1
		}
		if temp2 < 0 {
			temp2 = temp2 * -1
		}

		if temp1 < temp2 {
			stocks[i].Action = "sell"
			if sellValue < temp1 {
				sellValue = temp1
				if stocks[i].Date.After(buyDate) {
					sellDate = stocks[i].Date
				}
			}
		} else {
			stocks[i].Action = "buy"
			if buyValue > temp2 {
				buyValue = temp2
				buyDate = stocks[i].Date
			}
		}
	}

	return &stocks, buyDate.Format("02/01/2006"), sellDate.Format("02/01/2006"), err

}
