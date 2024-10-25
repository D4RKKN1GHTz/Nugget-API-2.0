package entity

import (
	"time"
)

type UserEntity struct {
	ID        string    `bun:"default:gen_ramdom_uuid(),pk"`
	FirstName string    `bun:"firstname,notnull"`
	LastName  string    `bun:"lastname,notnull"`
	Username  string    `bun:"username,notnull"`
	Password  string    `bun:"password,notnull"`
	CreatedAt int64     `bun:"created_at,nullzero,default:EXTRACT(EPOCH FROM NOW())"`
	UpdatedAt int64     `bun:"updated_at,nullzero,default:EXTRACT(EPOCH FROM NOW())"`
	DeletedAt time.Time `bun:"deleted_at,soft_delete"`
}
