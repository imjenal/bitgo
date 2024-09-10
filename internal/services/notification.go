package services

import (
	"bitgo/internal/models"
	"bitgo/internal/repository"
	"errors"
	"fmt"
)

type NotificationService interface {
	CreateNotification(notification *models.Notification) (int, error)
	ListNotifications(status string) ([]models.Notification, error)
	UpdateNotification(notification *models.Notification) error
	DeleteNotification(id int) error
}

type DefaultNotificationService struct {
	Repo repository.NotificationRepository
}

func NewNotificationService(repo repository.NotificationRepository) *DefaultNotificationService {
	return &DefaultNotificationService{Repo: repo}
}

func (s *DefaultNotificationService) CreateNotification(notification *models.Notification) (int, error) {
	if notification.CurrentPrice <= 0 {
		return 0, errors.New("current price must be greater than 0")
	}
	if notification.DailyChangePercentage == 0 {
		return 0, errors.New("daily change percentage can't be 0")
	}
	if len(notification.Emails) == 0 {
		return 0, errors.New("no emails provided")
	}
	id, err := s.Repo.Create(notification)
	if err != nil {
		return 0, fmt.Errorf("failed to create notification : %w", err)
	}
	return id, nil
}

func (s *DefaultNotificationService) ListNotifications(status string) ([]models.Notification, error) {
	if status == "" {
		return nil, errors.New("status can't be empty")
	}

	notifications, err := s.Repo.GetByStatus(status)
	if err != nil {
		return nil, fmt.Errorf("failed to list notifications : %w", err)
	}

	return notifications, nil
}

func (s *DefaultNotificationService) UpdateNotification(notification *models.Notification) error {
	if notification.ID <= 0 {
		return errors.New("invalid notification ID")
	}
	if notification.CurrentPrice <= 0 {
		return errors.New("current price must be greater than 0")
	}
	if notification.DailyChangePercentage == 0 {
		return errors.New("daily change percentage can't be 0")
	}
	if len(notification.Emails) == 0 {
		return errors.New("no emails provided")
	}

	err := s.Repo.Update(notification)
	if err != nil {
		return fmt.Errorf("failed to update notifications : %w", err)
	}

	return nil
}

func (s *DefaultNotificationService) DeleteNotification(id int) error {
	if id <= 0 {
		return errors.New("invalid notification ID")
	}
	err := s.Repo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete notifications : %w", err)
	}
	return nil
}
