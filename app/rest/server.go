package rest

import (
	"fmt"
	"log"

	_ "github.com/c-4u/loyalty-card/app/rest/docs"
	"github.com/c-4u/loyalty-card/domain/service"
	"github.com/c-4u/loyalty-card/infra/db"
	"github.com/c-4u/loyalty-card/infra/repo"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Loyalty Card Swagger API
// @version 1.0
// @description Swagger API for Loyalty Card Service.
// @termsOfService http://swagger.io/terms/

// @contact.name Coding4u
// @contact.email oi@coding4u.com.br

// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func StartRestServer(orm *db.DbOrm, port int) {
	r := fiber.New()
	r.Use(cors.New())

	repository := repo.NewRepository(orm)
	service := service.NewService(repository)
	restService := NewRestService(service)

	api := r.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/swagger/*", fiberSwagger.WrapHandler)
	{
		guest := v1.Group("/guests")
		guest.Post("", restService.CreateGuest)
		guest.Get("/:guest_id", restService.FindGuest)

		score := v1.Group("/scores")
		score.Post("", restService.CreateScore)
		score.Get("/:score_id", restService.FindScore)
		score.Post("/:score_id/use", restService.UseScore)
	}

	addr := fmt.Sprintf("0.0.0.0:%d", port)
	err := r.Listen(addr)
	if err != nil {
		log.Fatal("cannot start rest server", err)
	}

	log.Printf("rest server has been started on port %d", port)
}
