package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	project := os.Getenv("PUB_PROJECT")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/add-school", handler(project, "add-school"))
	e.GET("/add-grade", handler(project, "add-grade"))
	e.GET("/add-class", handler(project, "add-class"))
	e.GET("/add-teacher", handler(project, "add-teacher"))
	e.GET("/add-student", handler(project, "add-student"))

	e.Logger.Fatal(e.Start(":8080"))
}

func handler(project, path string) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		client, err := firestore.NewClient(ctx, project)
		if err != nil {
			log.Fatal(err)
		}
		defer client.Close()

		order := fmt.Sprintf("%s:%s", path, createUUID())

		_, err = client.Collection("operation").Doc(order).
			Set(ctx, map[string]interface{}{
				"order":    order,
				"sequence": time.Now().UnixNano(),
			}, firestore.MergeAll)
		if err != nil {
			log.Fatal(err)
		}
		return c.String(http.StatusOK, order)
	}
}

func createUUID() string {
	u, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}
	return u.String()
}
