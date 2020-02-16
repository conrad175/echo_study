package main

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo"
	"github.com/gomodule/redigo/redis"
)


func Connection() redis.Conn {
	const Addr = "redis:6379"
	c, err := redis.Dial("tcp", Addr)
	if err != nil {
		panic(err)
	}
	return c
}

func Set(key, value string, c redis.Conn) string{
	res, err := redis.String(c.Do("SET", key, value))
	if err != nil {
		panic(err)
	}
	return res
}

func main() {
	c := Connection()
	defer c.Close()
	res_set := Set("sample-key", "sample-value", c)
	fmt.Println(res_set) 

	e := echo.New()
	e.GET("/sample", func(c echo.Context) error {
	  return c.String(http.StatusOK, "echo sample !\n")
	})
	e.Logger.Fatal(e.Start(":80"))
}
