package user

import "time"

type User struct {
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	Grade     int       `json:"grade"`
	IsBlocked bool      `json:"is_blocked"`
	CreatedAt time.Time `json:"created_at"`
	Roles     []string  `json:"roles"`
}
