package utils

import "errors"

var (
	ErrEmailExists         = errors.New("Email already exists")
	ErrFillAllFields       = errors.New("Please fill in all required fields")
	ErrInvalidEmailFormat  = errors.New("Invalid email format")
	ErrInvalidCredentials  = errors.New("Invalid email or password")
	ErrUserNotFound        = errors.New("User not found")
	ErrFailedToCreateToken = errors.New("Failed to create authentication token")
	ErrTokenExpired        = errors.New("Token expired")
	ErrInvalidToken        = errors.New("Invalid auth token")
	ErrMissingToken        = errors.New("Missing token")
	ErrUnauthorizedAccess  = errors.New("Unauthorized access")
	ErrInvalidMovieID      = errors.New("Invalid movie ID")
	ErrMissingMovieData    = errors.New("Missing movie data")
	ErrMovieNotFound       = errors.New("Movie not found")
)
