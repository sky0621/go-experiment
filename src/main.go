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

	e.GET("/add-school", schoolHandler(project, "add-school"))
	e.GET("/edit-school/:sid", schoolHandler(project, "edit-school"))

	e.GET("/add-grade/:sid", handler(project, "add-grade"))
	e.GET("/add-class/:sid", handler(project, "add-class"))

	e.GET("/add-teacher/:sid", handler(project, "add-teacher"))
	e.GET("/add-student/:sid", handler(project, "add-student"))

	e.Logger.Fatal(e.Start(":8080"))
}

func schoolHandler(project, path string) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		client, err := firestore.NewClient(ctx, project)
		if err != nil {
			log.Fatal(err)
		}
		defer client.Close()

		operationSequence := createOperationSequence()

		schoolID := c.Param("sid")
		if schoolID == "" {
			schoolID = createUUID()

			_, err = client.Collection("school").Doc(schoolID).
				Set(ctx, map[string]interface{}{
					"state": "ACTIVE",
				}, firestore.MergeAll)
			if err != nil {
				log.Fatal(err)
			}
		}

		_, err = client.Collection("school").Doc(schoolID).
			Collection("operation").Doc(fmt.Sprintf("%d", operationSequence)).
			Set(ctx, map[string]interface{}{
				"operationSequence": operationSequence,
				"order":             path + ":" + schoolID,
			}, firestore.MergeAll)
		if err != nil {
			log.Fatal(err)
		}

		return c.String(http.StatusOK, schoolID)
	}
}

func handler(project, path string) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		client, err := firestore.NewClient(ctx, project)
		if err != nil {
			log.Fatal(err)
		}
		defer client.Close()

		schoolID := c.Param("sid")
		if schoolID == "" {
			log.Fatal("no schoolID")
		}

		newID := createUUID()

		operationSequence := createOperationSequence()

		_, err = client.Collection("school").Doc(schoolID).
			Collection("operation").Doc(fmt.Sprintf("%d", operationSequence)).
			Set(ctx, map[string]interface{}{
				"operationSequence": operationSequence,
				"order":             path + ":" + newID,
			}, firestore.MergeAll)
		if err != nil {
			log.Fatal(err)
		}
		return c.String(http.StatusOK, newID)
	}
}

func createUUID() string {
	u, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}
	return u.String()
}

func createOperationSequence() int64 {
	return time.Now().UnixNano()
}
