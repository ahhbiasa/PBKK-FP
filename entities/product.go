package entities

import (
	"time"
)

type Product struct {
	Id          int
	Name        string
	Category    Category
	Stock       int
	Description string
	Created_At  time.Time
	Updated_At  time.Time
}
