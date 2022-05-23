package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetIPAddress() echo.HandlerFunc {
	return func(c echo.Context) error {
		resp, err := http.Get("https://ifconfig.me/all.json")
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		var ipAddress map[string]interface{}
		err = json.Unmarshal(body, &ipAddress)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, ipAddress)
	}
}
func main() {
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		log.Fatal("$PORT must be set")
	}
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", GetIPAddress())
	e.Logger.Fatal(e.Start("0.0.0.0:" + httpPort))
}
