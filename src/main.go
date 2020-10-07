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

	e.GET("/add-region", regionHandler())
	e.GET("/edit-region/:rid", regionHandler())

	e.GET("/add-school/:rid", schoolHandler(project, "add-school"))
	e.GET("/edit-school/:rid/:sid", schoolHandler(project, "edit-school"))

	e.GET("/add-grade/:rid/:sid", handler(project, "add-grade"))
	e.GET("/add-class/:rid/:sid", handler(project, "add-class"))

	e.GET("/add-teacher/:rid/:sid", handler(project, "add-teacher"))
	e.GET("/add-student/:rid/:sid", handler(project, "add-student"))

	e.Logger.Fatal(e.Start(":8080"))
}

func regionHandler() func(c echo.Context) error {
	return func(c echo.Context) error {
		regionID := c.Param("rid")
		if regionID == "" {
			regionID = createUUID()
		}
		return c.String(http.StatusOK, regionID)
	}
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

		regionID := c.Param("rid")
		if regionID == "" {
			log.Fatal("no regionID")
		}

		schoolID := c.Param("sid")
		if schoolID == "" {
			schoolID = createUUID()

			_, err = client.Collection("region").Doc(regionID).
				Collection("school").Doc(schoolID).
				Set(ctx, map[string]interface{}{
					"state": "UNSYNCED",
				}, firestore.MergeAll)
			if err != nil {
				log.Fatal(err)
			}
		}

		_, err = client.Collection("region").Doc(regionID).
			Collection("school").Doc(schoolID).
			Collection("operation").Doc(operationSequence).
			Set(ctx, map[string]interface{}{
				"operationSequence": operationSequence,
				"order":             path + ":" + regionID + ":" + schoolID,
				"state":             "UNSYNCED",
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

		regionID := c.Param("rid")
		if regionID == "" {
			log.Fatal("no regionID")
		}
		schoolID := c.Param("sid")
		if schoolID == "" {
			log.Fatal("no schoolID")
		}

		newID := createUUID()

		operationSequence := createOperationSequence()

		_, err = client.Collection("region").Doc(regionID).
			Collection("school").Doc(schoolID).
			Collection("operation").Doc(operationSequence).
			Set(ctx, map[string]interface{}{
				"operationSequence": operationSequence,
				"order":             path + ":" + regionID + ":" + schoolID + ":" + newID,
				"state":             "UNSYNCED",
			}, firestore.MergeAll)
		if err != nil {
			log.Fatal(err)
		}
		return c.String(http.StatusOK, regionID+":"+schoolID+":"+newID)
	}
}

func createUUID() string {
	u, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}
	return u.String()
}

func createOperationSequence() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
