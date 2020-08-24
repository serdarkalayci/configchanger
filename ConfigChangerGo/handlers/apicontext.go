package handlers

// KeyRating is a key used for the Rating object in the context
type KeyRating struct{}

// APIContext is the struct that has a logger and validation instance. It's the base for all handler functions
type APIContext struct {
}

// NewAPIContext returns a new APIContext handler with the given logger
func NewAPIContext() *APIContext {
	return &APIContext{}
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}
