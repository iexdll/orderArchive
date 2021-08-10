package mongoDB

import (
	"go.mongodb.org/mongo-driver/mongo"
	"orderArchive/store"
)

type Store struct {
	db              *mongo.Database
	orderRepository *OrderRepository
}

func New(db *mongo.Database) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Order() store.OrderRepository {

	if s.orderRepository != nil {
		return s.orderRepository
	}

	s.orderRepository = &OrderRepository{
		store: s,
	}

	return s.orderRepository
}
