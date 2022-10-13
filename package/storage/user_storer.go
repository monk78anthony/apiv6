package storage

import (
	"context"
	"time"
)

type User struct {
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	Grade     int       `json:"grade"`
	IsBlocked bool      `json:"is_blocked"`
	CreatedAt time.Time `json:"created_at"`
	Roles     []string  `json:"roles"`
}

type UserStorer interface {
	Insert(ctx context.Context, user User) error
	Find(ctx context.Context, uuid string) (User, error)
	Delete(ctx context.Context, uuid string) error
	Update(ctx context.Context, user User) error
}
