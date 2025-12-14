package errors

import pkgerrors "yoyoj-go-rag/pkg/errors"

// 用户域错误码前缀: USER
const (
	// 用户实体相关错误 (USER-E-xxx)
	ErrCodeInvalidUsername = "USER-E-001"
	ErrCodeInvalidEmail    = "USER-E-002"
	ErrCodeInvalidPassword = "USER-E-003"
	ErrCodeUserNotFound    = "USER-E-004"
	ErrCodeUserExists      = "USER-E-005"
	ErrCodePasswordHashFailed = "USER-E-006"

	// 用户业务逻辑错误 (USER-B-xxx)
	ErrCodeRegistrationFailed = "USER-B-001"
	ErrCodeLoginFailed        = "USER-B-002"
	ErrCodeUnauthorized       = "USER-B-003"
)

// 预定义错误
var (
	ErrInvalidUsername = pkgerrors.NewDomainError(
		ErrCodeInvalidUsername,
		"invalid username, must be a non-empty string",
	)

	ErrInvalidEmail = pkgerrors.NewDomainError(
		ErrCodeInvalidEmail,
		"invalid email format",
	)

	ErrInvalidPasswordLength = pkgerrors.NewDomainError(
		ErrCodeInvalidPassword,
		"invalid password, must be at least 6 characters long",
	)

	ErrInvalidPassword = pkgerrors.NewDomainError(
		ErrCodeInvalidPassword,
		"invalid password, must uppercase, lowercase letters and numbers",
	)

	ErrUserNotFound = pkgerrors.NewDomainError(
		ErrCodeUserNotFound,
		"user not found",
	)

	ErrUserExists = pkgerrors.NewDomainError(
		ErrCodeUserExists,
		"user already exists",
	)

	ErrPasswordHashFailed = pkgerrors.NewDomainError(
		ErrCodePasswordHashFailed,
		"failed to hash password",
	)
)
