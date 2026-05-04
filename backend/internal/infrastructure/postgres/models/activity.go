package models

import "time"

type Activity struct {
	ID                string    `db:"id"`
	ChildProfileID    int64     `db:"child_profile_id"`
	Type              string    `db:"type"`
	ActivityID        string    `db:"activity_id"`
	ConfirmedByParent bool      `db:"confirmed_by_parent"`
	CreatedAt         time.Time `db:"created_at"`
}
