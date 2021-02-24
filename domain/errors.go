package domain

// AppError AppError
type AppError struct {
	ErrType ErrorType
	Err     error
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

// ErrorType ErrorType
type ErrorType int

const (
	// ErrDatabase ErrorDatabase
	ErrDatabase ErrorType = iota
	// ErrValidation ErrValidation
	ErrValidation
	// ErrEmailEmpty ErrEmailEmpty
	ErrEmailEmpty
)
