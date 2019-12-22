package main

import (
  "net/http"
  "github.com/labstack/echo"
)

func main() {
  e := echo.New()
  e.GET("/sample", func(c echo.Context) error {
    return c.String(http.StatusOK, "echo sample !\n")
  })
  e.Logger.Fatal(e.Start(":80"))
}
