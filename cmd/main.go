package main

import (
	"bitgo/internal/db"
	"bitgo/internal/handlers"
	"bitgo/internal/repository"
	"bitgo/internal/services"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbConn, err := db.NewPostgresDB()
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	defer dbConn.Close()

	notificationRepo := repository.NewPostgresNotificationRepository(dbConn)
	notificationService := services.NewNotificationService(notificationRepo)
	notificationHandler := handlers.NewNotificationHandler(notificationService)

	r := mux.NewRouter()

	r.HandleFunc("/notifications", notificationHandler.CreateNotification).Methods(http.MethodPost)
	r.HandleFunc("/notifications", notificationHandler.ListNotifications).Methods(http.MethodGet)
	r.HandleFunc("/notifications/{id:[0-9]+}", notificationHandler.UpdateNotification).Methods(http.MethodPut)
	r.HandleFunc("/notifications/{id:[0-9]+}", notificationHandler.DeleteNotification).Methods(http.MethodDelete)

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)

}
