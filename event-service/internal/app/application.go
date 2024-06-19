package app

import (
	"event-service/config"
	"event-service/internal/domain/service"
	"event-service/internal/infra/db"
	"event-service/internal/infra/repository"
	"event-service/internal/interfaces/api"
	"event-service/pkg/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start(cfg config.Config) {
	dbConn := db.NewPostgresDB(cfg.DBSource)

	eventRepo := repository.NewEventRepository(dbConn)
	eventService := service.NewEventService(eventRepo)
	eventHandler := api.NewEventHandler(eventService)

	categoryRepo := repository.NewCategoryRepository(dbConn)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := api.NewCategoryHandler(categoryService)

	router := mux.NewRouter()

	// Http Logger
	router.Use(utils.LoggingMiddleware)

	// Event Service
	router.HandleFunc("/events", eventHandler.CreateEvent).Methods("POST")
	router.HandleFunc("/events", eventHandler.GetAllEvents).Methods("GET")
	router.HandleFunc("/events/{id:[0-9]+}", eventHandler.GetEventByID).Methods("GET")
	router.HandleFunc("/events/{id:[0-9]+}", eventHandler.UpdateEvent).Methods("PUT")
	router.HandleFunc("/events/{id:[0-9]+}", eventHandler.DeleteEvent).Methods("DELETE")

	// Category Service
	router.HandleFunc("/categories", categoryHandler.CreateCategory).Methods("POST")
	router.HandleFunc("/categories", categoryHandler.GetAllCategories).Methods("GET")
	router.HandleFunc("/categories/{id:[0-9]+}", categoryHandler.GetCategoryByID).Methods("GET")
	router.HandleFunc("/categories/{id:[0-9]+}", categoryHandler.UpdateCategory).Methods("PUT")
	router.HandleFunc("/categories/{id:[0-9]+}", categoryHandler.DeleteCategory).Methods("DELETE")

	log.Printf("Starting server at %s", cfg.ServerAddress)
	log.Fatal(http.ListenAndServe(cfg.ServerAddress, router))
}
