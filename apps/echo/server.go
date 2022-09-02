package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResponseType struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	e := echo.New()

	e.GET("/version", func(c echo.Context) error {
		response := &ResponseType{
			Status: http.StatusOK,
			Data:   "v1",
		}

		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusOK)

		return c.JSON(c.Response().Status, response)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
