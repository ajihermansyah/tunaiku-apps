package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

//** TYPES

type (
	Stock struct {
		ID      bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Date    time.Time     `bson:"date" json:"-" form:"date,01/02/2006"`
		Open    int           `bson:"open" json:"open" form:"open"`
		High    int           `bson:"high" json:"high" form:"high"`
		Low     int           `bson:"low" json:"low" form:"low"`
		Close   int           `bson:"close" json:"close" form:"close"`
		DateStr string        `bson:"-" json:"date_str"`
		Action  string        `bson:"-" json:"action"`
	}
)
