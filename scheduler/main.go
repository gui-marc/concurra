package main

import (
	"context"
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/gui-marc/concurra/repository"
	"github.com/gui-marc/concurra/scheduler/handlers"
	"github.com/gui-marc/concurra/scheduler/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Cli struct {
	Port        int    `help:"Port to run the server on" default:"8080"`
	DatabaseURL string `help:"Database URL" default:"postgres://concurra:concurra@localhost:5432/concurra"`

	CORS struct {
		AllowOrigins []string `help:"Allowed origins for CORS" default:"*"`
	} `embed:"" prefix:"cors."`
}

func main() {
	var cli Cli
	kong.Parse(&cli)

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: cli.CORS.AllowOrigins,
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	q, conn, err := repository.NewPGQueries(cli.DatabaseURL)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}
	defer conn.Close(context.Background())

	eventRepository := repository.NewPGXEventRepository(q)
	eventService := services.NewEventService(eventRepository)
	eventHandler := handlers.NewEventHandler(eventService)

	v1 := e.Group("/api/v1")

	v1.GET("/events", eventHandler.GetEvents)
	v1.POST("/events", eventHandler.CreateEvent)
	v1.GET("/event/:eventId", eventHandler.GetEventByID)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cli.Port)))
}
