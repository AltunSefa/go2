package models

import (
	"github.com/google/uuid"
)

type Book struct {
	ID        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string     `gorm:"type:varchar(100);not null"`
	UserID    *uuid.UUID `gorm:"type:uuid;not null"`
}