package entities

import (
	"time"
)

type Shop struct {
	Id        int
	Name      string
	Address   string
	Products  []Product
	CreatedAt time.Time
	UpdatedAt time.Time
}
