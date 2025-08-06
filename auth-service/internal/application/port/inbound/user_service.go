package inbound

type UserService interface {
	// CreateUser creates a new user with the given details.
	CreateUser(name string, email string, password string) (int64, error)
}
