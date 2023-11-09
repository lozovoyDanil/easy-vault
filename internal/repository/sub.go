package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
	"main.go/internal/model"
)

type SubRepository struct {
	db *sqlx.DB
}

func NewSubRepository(db *sqlx.DB) *SubRepository {
	return &SubRepository{db: db}
}

func (r *SubRepository) SubByUserId(id int) (model.Subscription, error) {
	return model.Subscription{
		Id:        1,
		Tier:      1000,
		ExpiresAt: time.Now().Add(1 * time.Hour),
	}, nil
}

func (r *SubRepository) CreateSub(input model.Subscription) error {
	return nil
}

func (r *SubRepository) UpdateSub(input model.Subscription) error {
	return nil
}
