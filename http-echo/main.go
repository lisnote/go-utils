package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"libs/version"
)

func main() {
	version.ShowVersionDetect("1.0.0")
	if len(os.Args) < 2 {
		log.Fatal("usage: http-echo <host:port>")
	}
	addr := os.Args[1]
	e := echo.New()

	e.Any("/*", func(c echo.Context) error {
		// 读取请求体
		var body map[string]any
		if err := c.Bind(&body); err != nil {
			body = map[string]any{}
		}

		// 构建返回 JSON
		resp := map[string]any{
			"method":    c.Request().Method,
			"host":      c.Request().Host,
			"path":      c.Path(),
			"query":     c.QueryParams(),
			"headers":   c.Request().Header,
			"data":      body,
			"ipaddress": c.RealIP(),
		}

		return c.JSON(http.StatusOK, resp)
	})

	e.Logger.Fatal(e.Start(addr))
}
