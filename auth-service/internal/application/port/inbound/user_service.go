package inbound

import "gin/internal/interface/dto/http"

type UserService interface {
	CreateUser(userRequest http.UserCreateRequest) (http.UserDetailsResponse, error)
}
