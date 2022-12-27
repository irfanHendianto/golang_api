package helper

import "strings"

// Struct Response used for return response
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

// Struct empty oject used for when data doesn't want to be null on response
type EmptyObject struct{}

// Function for return response success
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Error:   nil,
		Data:    data,
	}
	return res
}

// Function for return resoinse error
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Error:   splittedError,
		Data:    data,
	}

	return res
}

// Function for check password when register / update
func ComparedPassword(password string, confirmPassword string) bool {
	return (password == confirmPassword)
}
