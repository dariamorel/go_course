package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"project/accounts"
)

func main() {
	accountsHandler := accounts.New()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/account", accountsHandler.GetAccount)
	e.POST("/account/create", accountsHandler.CreateAccount)
	e.POST("account/patch", accountsHandler.PatchAccount)
	e.POST("account/change", accountsHandler.ChangeAccount)
	e.POST("account/delete", accountsHandler.DeleteAccount)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
