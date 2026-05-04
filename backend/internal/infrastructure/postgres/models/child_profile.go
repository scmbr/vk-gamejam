package models

import (
	"database/sql"
	"time"
)

type ChildProfileModel struct {
	ID            int64          `db:"id"`
	UserID        int64          `db:"user_id"`
	Name          string         `db:"child_name"`
	Gender        string         `db:"child_gender"`
	ParentPin     sql.NullString `db:"parent_pin"`
	HasPet        bool           `db:"has_pet"`
	IsFirstLaunch bool           `db:"is_first_launch"`
	LastLogin     time.Time      `db:"last_login"`
	LastLogout    time.Time      `db:"last_logout"`
}
