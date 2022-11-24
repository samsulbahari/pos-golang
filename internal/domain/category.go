package domain

import (
	"time"

	"gorm.io/gorm"
)

type MCategory struct {
	ID        int
	Name      string `json:"name" form:"name" binding:"required"`
	Image     string
	Status    *bool `form:"status" binding:"required,boolean"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-"`
}

type Update struct {
	ID        int
	Name      string `json:"name" form:"name"`
	Image     string
	Status    *bool `form:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-"`
}

type ResultCategory struct {
	Total    int
	PerPage  int
	Page     int
	LastPage float64
	Data     []MCategory
}
