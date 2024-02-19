package dto

type CategoryInput struct {
	Title string `json:"title"`
	UserID string `json:"user_id"`
}

type UserOwnerCategory struct {
	ID string `json:"id"`
	Firstname string `json:"firstname"`
}

type CategoryUserResponse struct {
	ID string `json:"id"`
	Title string `json:"title"`
	User UserOwnerCategory `json:"user"`
}
