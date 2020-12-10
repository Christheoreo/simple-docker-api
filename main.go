package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/christheoreo/docker-test-api/prime"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type res struct {
	Message string `json:"message"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/:number", func(c echo.Context) error {
		numberS := c.Param("number")
		number, err := strconv.Atoi(numberS)
		if err != nil {
			return c.JSONPretty(http.StatusBadRequest, res{
				Message: "The number provided could not be converted from a string to an int",
			}, "  ")
		}
		p := prime.IsNumberPrime(int64(number))
		r := &res{
			Message: "NA",
		}
		if p {
			r.Message = fmt.Sprintf("%d is a prime number!", number)
		} else {
			r.Message = fmt.Sprintf("%d is not a prime number!", number)
		}
		return c.JSONPretty(http.StatusOK, r, "  ")
	})
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
