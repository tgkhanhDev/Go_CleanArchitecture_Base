package inbound

type AuthService interface {
	Login(username, password string) (string, error)
	//Introspect(token string) (map[string]interface{}, error);
	//Refresh(token string) (string, error);
	//Logout(token string) error;
}
