package services

import authReq "gin/internal/application/dtos/request"

type AuthService interface {
	Login(username, password string) (string, error)
	RegisterUser(registerReq authReq.RegisterRequest)
	//Introspect(token string) (map[string]interface{}, error);
	//Refresh(token string) (string, error);
	//Logout(token string) error;
}
