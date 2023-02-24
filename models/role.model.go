package models

import ()

type Role struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	Name      string     `gorm:"type:varchar(100);not null"`
}