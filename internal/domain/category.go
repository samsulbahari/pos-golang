package domain

import "time"

type MCategory struct {
	ID        int
	Name      string
	Image     string
	Status    bool
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
