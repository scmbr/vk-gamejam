package dto

import "time"

type UserMeResponse struct {
	ID              int64     `json:"id"`
	Email           string    `json:"email"`
	CreatedAt       time.Time `json:"createdAt"`
	HasChildProfile bool      `json:"hasChildProfile"`
}
