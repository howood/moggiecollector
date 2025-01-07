package response

type UserResponse struct {
	UserID uint64 `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status int64  `json:"status"`
}
