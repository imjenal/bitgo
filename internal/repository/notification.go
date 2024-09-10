package repository

import (
	"bitgo/internal/models"
	"database/sql"
	"errors"
)

type NotificationRepository interface {
	Create(notification *models.Notification) (int, error)
	GetByStatus(status string) ([]models.Notification, error)
	Update(notification *models.Notification) error
	Delete(id int) error
}

type PostgresNotificationRepository struct {
	DB *sql.DB
}

func NewPostgresNotificationRepository(db *sql.DB) *PostgresNotificationRepository {
	return &PostgresNotificationRepository{DB: db}
}

func (r *PostgresNotificationRepository) Create(notification *models.Notification) (int, error) {
	query := `INSERT INTO notifications (current_price, daily_change_percentage, trading_volume, emails, status) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id int
	err := r.DB.QueryRow(query, notification.CurrentPrice, notification.DailyChangePercentage, notification.TradingVolume, notification.Emails, notification.Status).Scan(&id)
	return id, err
}

func (r *PostgresNotificationRepository) GetByStatus(status string) ([]models.Notification, error) {
	query := `SELECT id, current_price, daily_change_percentage, trading_volume, emails, status, created_at, updated_at FROM notifications WHERE status = $1`
	rows, err := r.DB.Query(query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []models.Notification
	for rows.Next() {
		var notification models.Notification
		if err := rows.Scan(&notification.ID, &notification.CurrentPrice,
			&notification.DailyChangePercentage, &notification.TradingVolume, &notification.Emails,
			&notification.Status, &notification.CreatedAt, &notification.UpdatedAt); err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return notifications, nil
}

func (r *PostgresNotificationRepository) Update(notification *models.Notification) error {
	query := `UPDATE notifications SET current_price = $1, daily_change_percentage = $2, trading_volume =$3, emails = $4, status =$5 updated_at = NOW() WHERE id =$6`
	result, err := r.DB.Exec(query, notification.CurrentPrice, notification.DailyChangePercentage, notification.TradingVolume, notification.Emails,
		notification.Status, notification.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows affected, notifications not found")
	}
	return nil
}

func (r *PostgresNotificationRepository) Delete(id int) error {
	query := `DELETE FROM notifications WHERE id =$1`
	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows affected, notifications not found")
	}
	return nil
}
