package models

type ChildProfileModel struct {
	ID        int64  `db:"id"`
	UserID    int64  `db:"user_id"`
	Name      string `db:"name"`
	Gender    string `db:"gender"`
	ParentPin string `db:"parent_pin"`
	HasPet    bool   `db:"has_pet"`
}