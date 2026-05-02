package domain

import "time"

type Pet struct {
	UserID     int64
	Name       string
	Type       string
	Gender     string
	Level      int
	XP         float64
	Knowledge  float64
	Energy     float64
	Creativity float64
	Sport      float64
	LastOnline time.Time
}