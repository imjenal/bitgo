package handlers

import (
	"bitgo/internal/models"
	"bitgo/internal/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type NotificationHandler struct {
	Service services.NotificationService
}

func NewNotificationHandler(service services.NotificationService) *NotificationHandler {
	return &NotificationHandler{Service: service}
}
func (h *NotificationHandler) CreateNotification(w http.ResponseWriter, r *http.Request) {
	var notification models.Notification
	if err := json.NewDecoder(r.Body).Decode(&notification); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.Service.CreateNotification(&notification)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":     id,
		"status": "Created",
	})

}

func (h *NotificationHandler) ListNotifications(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	notifications, err := h.Service.ListNotifications(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(notifications)
}

func (h *NotificationHandler) UpdateNotification(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid notification ID", http.StatusBadRequest)
		return
	}

	var notification models.Notification
	if err := json.NewDecoder(r.Body).Decode(&notification); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	notification.ID = id

	err = h.Service.UpdateNotification(&notification)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"status": "Updated"})

}

func (h *NotificationHandler) DeleteNotification(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid notification ID", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteNotification(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"status": "Deleted"})

}
