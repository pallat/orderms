package store

import (
	"github.com/pallat/micro/order"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormStore struct {
	db *gorm.DB
}

func NewMariaDBStore(dsn string) *GormStore {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&order.Order{})

	return &GormStore{db: db}
}

func (s *GormStore) Save(o order.Order) error {
	return s.db.Create(o).Error
}
