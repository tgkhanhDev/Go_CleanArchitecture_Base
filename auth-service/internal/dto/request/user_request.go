package http

type UserCreateRequest struct {
	Username string
	Password string
	Email    string
}

func CreateUserRequest(username, password, email string) map[string]string {
	return map[string]string{
		"username": username,
		"password": password,
		"email":    email,
	}
}
