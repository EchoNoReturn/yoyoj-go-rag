package errors

import "fmt"

type DomainError interface {
	error
	Code() string
	Message() string
	Details() map[string]interface{}
}

type BaseDomainError struct {
	code    string
	message string
	details map[string]interface{}
}

func NewDomainError(code, message string) *BaseDomainError {
	return &BaseDomainError{
		code:    code,
		message: message,
		details: make(map[string]interface{}),
	}
}

func (e *BaseDomainError) Error() string {
	return fmt.Sprintf("[%s] %s", e.code, e.message)
}

func (e *BaseDomainError) Code() string {
	return e.code
}

func (e *BaseDomainError) Message() string {
	return e.message
}

func (e *BaseDomainError) Details() map[string]interface{} {
	return e.details
}

func (e *BaseDomainError) WithDetail(key string, value interface{}) *BaseDomainError {
	e.details[key] = value
	return e
}
