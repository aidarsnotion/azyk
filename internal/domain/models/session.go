package models

import "time"

type Session struct {
	ID        int       `json:"id"`
	UserId    int       `json:"user_id"`
	DeviceId  string    `json:"device_id"`
	CreatedAt time.Time `json:"created"`
}
