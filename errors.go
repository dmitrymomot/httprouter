package httprouter

import "errors"

// Predefined errors.
var (
	ErrFailedToEncodeResponse = errors.New("failed to encode response")
	ErrResponsePayloadIsNil   = errors.New("response payload is nil")
	ErrResponseMustBeString   = errors.New("response must be a string or byte slice")
)
