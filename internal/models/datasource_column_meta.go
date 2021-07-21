package models

import (
	"time"

	"gorm.io/gorm"
)

type DatasourceColumnMeta struct {
	ID          uint64 `gorm:"primaryKey"`
	ColumnId    uint64 `gorm:"index;not null"`
	Cardinarity uint32 `gorm:"default:0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
