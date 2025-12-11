package apperror

import "net/http"

// AppError representa un error de aplicación con código HTTP, mensaje y estado.
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

// Error implementa la interfaz error.
func (e AppError) Error() string { return e.Message }

// FromCode crea un AppError usando el texto estándar del status HTTP.
func FromCode(code int, message string) AppError {
	return AppError{Code: code, Message: message, Status: http.StatusText(code)}
}
