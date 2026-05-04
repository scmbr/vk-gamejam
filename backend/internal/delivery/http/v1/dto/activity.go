package dto

type CreateActivityRequest struct {
	Type              string `json:"type" binding:"required"`
	ActivityID        string `json:"activityId" binding:"required"`
	ConfirmedByParent bool   `json:"confirmedByParent"`
}

type ActivityResponse struct {
	ID                string `json:"id"`
	Type              string `json:"type"`
	ActivityID        string `json:"activityId"`
	ConfirmedByParent bool   `json:"confirmedByParent"`
	CreatedAt         string `json:"createdAt"`
}
