package http

func CreateUserRequest(username, password, email string) map[string]string {
	return map[string]string{
		"username": username,
		"password": password,
		"email":    email,
	}
}
