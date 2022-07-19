package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net"
	"net/http"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {

		addrs, err := net.InterfaceAddrs()
		if err != nil {
			panic(err)
		}
		html := ""

		for _, address := range addrs {
			if ipnet, ok := address.(*net.IPNet); ok {
				if ipnet.IP.To4() != nil {
					ipAddr := ipnet.IP.String()
					html += ipAddr + "<br>"
				}
			}
		}
		return c.HTML(http.StatusOK, html)
	})

	e.Logger.Fatal(e.Start(":80"))
}
