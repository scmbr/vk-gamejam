package domain

type ChildProfile struct {
	ID        int64
	UserID    int64
	Name      string
	Gender    string
	ParentPin string
	HasPet    bool
}