// Copyright 2013 Ardan Studios. All rights reserved.
// Use of baseController source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

// Package baseController implements boilerplate code for all baseControllers.
package base

import (
	"github.com/astaxie/beego"
	"tunaiku/services"
	"tunaiku/utilities/mongo"
)

//** TYPES

type (
	// BaseController composes all required types and behavior.
	BaseController struct {
		beego.Controller
		services.Service
	}
)

//** INTERCEPT FUNCTIONS

// Prepare is called prior to the baseController method.
func (baseController *BaseController) Prepare() {
	beego.SetStaticPath("/uploads", "tmp/uploads")
	beego.ReadFromRequest(&baseController.Controller)

	baseController.UserID = baseController.GetString("userID")

	if baseController.UserID == "" {
		baseController.UserID = baseController.GetString(":userID")
	}

	if baseController.UserID == "" {
		baseController.UserID = "Unknown"
	}

	if err := baseController.Service.Prepare(); err != nil {
		baseController.ServeError(err)
		return
	}
}

// Finish is called once the baseController method completes.
func (baseController *BaseController) Finish() {
	defer func() {
		if baseController.MongoSession != nil {
			mongo.CloseSession(baseController.UserID, baseController.MongoSession)
			baseController.MongoSession = nil
		}
	}()
}

//** EXCEPTIONS

// ServeError prepares and serves an Error exception.
func (baseController *BaseController) ServeError(err error) {
	baseController.Data["json"] = struct {
		Error string `json:"Error"`
	}{err.Error()}
	baseController.Ctx.Output.SetStatus(500)
	baseController.ServeJSON()
}
