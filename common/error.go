package common

import "fmt"

type BadRequestError struct {
	Message string
	Details map[string]interface{} // Optional details
}

func (e *BadRequestError) Error() string {
	return fmt.Sprintf("Bad Request: %s", e.Message)
}

type ServerError struct {
	Message string
	Err     error // Embed the original error
}

func (e *ServerError) Error() string {
	return fmt.Sprintf("Server Error: %s. Original error: %v", e.Message, e.Err)
}

type DatabaseError struct {
	Message string
	Err     error
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("Database Error: %s. Original error: %v", e.Message, e.Err)
}

type CustomError struct {
	Message string
	Code    int
	Err     error
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Custom Error (Code %d): %s. Original error: %v", e.Code, e.Message, e.Err)
}

func NewBadRequestError(message string, details map[string]interface{}) *BadRequestError {
	return &BadRequestError{
		Message: message,
		Details: details,
	}
}

func NewServerError(message string, err error) *ServerError {
	return &ServerError{
		Message: message,
		Err:     err,
	}
}

func NewDatabaseError(message string, err error) *DatabaseError {
	return &DatabaseError{
		Message: message,
		Err:     err,
	}
}

func NewCustomError(message string, code int, err error) *CustomError {
	return &CustomError{
		Message: message,
		Code:    code,
		Err:     err,
	}
}
