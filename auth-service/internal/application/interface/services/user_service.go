package services

import (
	req "gin/internal.bak/dto/request"
	res "gin/internal.bak/dto/response"
)

type UserService interface {
	CreateUser(userRequest req.UserCreateRequest) (res.UserDetailsResponse, error)
}
