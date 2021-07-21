package models

import (
	"time"

	"gorm.io/gorm"
)

type DatasourceColumn struct {
	ID           uint64 `gorm:"primaryKey"`
	DatasourceId uint64 `gorm:"index;not null"`
	ColumnName   string `gorm:"size:255"`
	ColumnType   string `gorm:"type:enum('dimension', 'metric');default:'dimension'"`
	IsCount      bool   `gorm:"default:false"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
