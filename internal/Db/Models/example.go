package Models

import (
	"time"
)

type Example struct {
	Id        []uint8    `db:"id" json:"id"`
	Sub       string     `db:"sub" json:"sub"`
	Label     string     `db:"label" json:"label"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}
