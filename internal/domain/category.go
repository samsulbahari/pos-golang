package domain

import "time"

type MCategory struct {
	ID        int
	Name      string `json:"name" form:"name" binding:"required"`
	Image     string
	Status    bool `form:"status" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ResultCategory struct {
	Total    int
	PerPage  int
	Page     int
	LastPage float64
	Data     []MCategory
}
