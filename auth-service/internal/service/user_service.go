package inbound

import (
	req "gin/internal/dto/request"
	res "gin/internal/dto/response"
)

type UserService interface {
	CreateUser(userRequest req.UserCreateRequest) (res.UserDetailsResponse, error)
}
