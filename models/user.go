package model

import (
	"time"
)

// User 用户表
var User struct {
	ID        int        `db:"id"`
	Name      string     `db:"name"`
	Telephone int        `db:"telephone"`
	Avatar    string     `db:"avatar"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"delete_at"`
}
