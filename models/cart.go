package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	User  string `json:"user"`
	Total int    `json:"total"`
}
