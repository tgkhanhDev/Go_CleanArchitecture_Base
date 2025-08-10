package dto

type APIResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}

func JSONResponse(data interface{}, message string, code int) APIResponse {
	return APIResponse{
		Data:    data,
		Message: message,
		Code:    code,
	}
}

// Some useful constants
// 200 OK
func SuccessResponse(data interface{}, message string) APIResponse {
	return JSONResponse(data, message, 200)
}

// 201 Created
func CreatedResponse(data interface{}, message string) APIResponse {
	return JSONResponse(data, message, 201)
}

// 400 Bad Request
func BadRequestResponse(message string) APIResponse {
	return JSONResponse(nil, message, 400)
}

// 401 Unauthorized
func UnauthorizedResposne(message string) APIResponse {
	return JSONResponse(nil, message, 401)
}

// 404 Not Found
func NotFoundResponse(message string) APIResponse {
	return JSONResponse(nil, message, 404)
}

func NotFoundWithDataResponse(data interface{}, message string) APIResponse {
	return JSONResponse(data, message, 404)
}
