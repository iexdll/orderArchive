package store

import (
	"orderArchive/models"
	"time"
)

type OrderRepository interface {
	Save(o *models.Order) error
	Get(string) (*models.Order, error)
	FindIDByDate(time.Time, time.Time) ([]string, error)
	FindIDByCustomer(string, string, int32, int32) ([]string, error)
}
