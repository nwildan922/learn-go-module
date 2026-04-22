package repository

import (
	"context"

	"github.com/nwildan922/learn-go-module/model"
	"github.com/nwildan922/learn-go-module/pkg/db"
	"gorm.io/gorm"
)

type CounterRepository struct {
	db *gorm.DB
}

func NewCounterRepository(database *db.Database) *CounterRepository {
	return &CounterRepository{
		db: database.DB,
	}
}

// Create counter record
func (r *CounterRepository) Create(ctx context.Context, data *model.Counter) error {
	return r.db.WithContext(ctx).Create(data).Error
}
