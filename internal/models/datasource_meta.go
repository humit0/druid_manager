package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type DatasourceMeta struct {
	ID           uint64         `gorm:"primaryKey"`
	DatasourceId uint64         `gorm:"index;not null"`
	DateId       datatypes.Date `gorm:"index"`
	RowCnt       uint64
	SegmentCnt   uint64
	ByteSize     uint64
	RollupRatio  float64 `gorm:"type:decimal(10,2);"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
