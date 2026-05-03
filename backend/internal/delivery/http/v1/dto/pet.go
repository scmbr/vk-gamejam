package dto

import "time"

type PetStateDTO struct {
	PetName    string    `json:"petName"`
	PetType    string    `json:"petType"`
	PetGender  string    `json:"petGender"`
	PetLevel   int       `json:"petLevel"`
	CurrentXP  float64   `json:"currentXP"`
	Knowledge  float64   `json:"knowledge"`
	Energy     float64   `json:"energy"`
	Creativity float64   `json:"creativity"`
	Sport      float64   `json:"sport"`
	LastOnline time.Time `json:"lastOnline"`
}

type CreatePetRequest struct {
	PetName   string `json:"petName" binding:"required"`
	PetType   string `json:"petType" binding:"required"`
	PetGender string `json:"petGender" binding:"required"`
}
type UpdatePetStateRequest struct {
	PetLevel   int       `json:"petLevel"`
	CurrentXP  float64   `json:"currentXP"`
	Knowledge  float64   `json:"knowledge"`
	Energy     float64   `json:"energy"`
	Creativity float64   `json:"creativity"`
	Sport      float64   `json:"sport"`
	LastOnline time.Time `json:"lastOnline"`
}
type PetResponse struct {
	PetName    string    `json:"petName"`
	PetType    string    `json:"petType"`
	PetGender  string    `json:"petGender"`
	PetLevel   int       `json:"petLevel"`
	CurrentXP  float64   `json:"currentXP"`
	Knowledge  float64   `json:"knowledge"`
	Energy     float64   `json:"energy"`
	Creativity float64   `json:"creativity"`
	Sport      float64   `json:"sport"`
	LastOnline time.Time `json:"lastOnline"`
}
