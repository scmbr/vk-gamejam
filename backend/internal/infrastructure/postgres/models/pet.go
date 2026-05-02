package models

import "time"

type PetModel struct {
	UserID     int64     `db:"user_id"`
	Name       string    `db:"name"`
	Type       string    `db:"type"`
	Gender     string    `db:"gender"`
	Level      int       `db:"level"`
	XP         float64   `db:"xp"`
	Knowledge  float64   `db:"knowledge"`
	Energy     float64   `db:"energy"`
	Creativity float64   `db:"creativity"`
	Sport      float64   `db:"sport"`
	LastOnline time.Time `db:"last_online"`
}