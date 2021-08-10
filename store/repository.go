package store

import (
	"orderArchive/models"
	"time"
)

type OrderRepository interface {
	Save(o *models.Order) error
	Get(string) (*models.Order, error)
	FindByDate(time.Time, time.Time) ([]string, error)
}
