package domainerror

type DomainError struct {
	Message string `json:"message"`
}

// Error treatment
func New(message string) *DomainError {
	return &DomainError{
		Message: message,
	}
}
