package domain

import "time"

type ChildProfile struct {
	ID            int64
	UserID        int64
	Name          string
	Gender        string
	ParentPin     *string
	HasPet        bool
	IsFirstLaunch bool
	LastLogin     time.Time
	LastLogout    time.Time
}
