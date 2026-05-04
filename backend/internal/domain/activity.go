package domain

import "time"

type ActivityType string

const (
	ActivityReading ActivityType = "reading"
	ActivityArt     ActivityType = "art"
	ActivitySport   ActivityType = "sport"
)

type Activity struct {
	ID                string
	ChildProfileID    int64
	Type              ActivityType
	ActivityID        string
	ConfirmedByParent bool
	CreatedAt         time.Time
}
