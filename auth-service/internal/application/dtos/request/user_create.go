package request

type CreateAccountRequest struct {
	Email        string `json:"email" validate:"required,email,max=255"`
	PasswordHash string `json:"password_hash" validate:"required,min=8,max=255"`
}
