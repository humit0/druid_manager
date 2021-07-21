package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Datasource struct {
	ID        uint64         `gorm:"primaryKey"`
	Name      string         `gorm:"size:255;unique;not null"`
	Owner     sql.NullString `gorm:"size:80"`
	Interval  uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
