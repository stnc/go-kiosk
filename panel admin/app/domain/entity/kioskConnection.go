package entity

import (
	"time"
)

//KioskConnection strcut
type KioskConnection struct {
	ID          uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Title       string     `gorm:"size:255 ;not null;" validate:"required" json:"title" `
	Content     string     `gorm:"type:text ;" validate:"required"  json:"content"`
	StructureNo int        `gorm:"type:smallint ;NOT NULL;DEFAULT:'1'" validate:"required"`
	Status      int        `gorm:"type:smallint ;NOT NULL;DEFAULT:'1'" validate:"required"`
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
