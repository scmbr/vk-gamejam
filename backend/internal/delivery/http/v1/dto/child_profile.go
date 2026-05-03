package dto

type ChildProfileResponse struct {
	Name          string  `json:"name"`
	Gender        string  `json:"gender"`
	ParentPin     *string `json:"parentPin,omitempty"`
	HasPet        bool    `json:"hasPet"`
	IsFirstLaunch bool    `json:"isFirstLaunch"`
}
type CreateChildProfileRequest struct {
	Name      string  `json:"name" binding:"required"`
	Gender    string  `json:"gender" binding:"required"`
	ParentPin *string `json:"parentPin"`
}
type UpdateChildProfileRequest struct {
	Name      *string `json:"name"`
	Gender    *string `json:"gender"`
	ParentPin *string `json:"parentPin"`
}
