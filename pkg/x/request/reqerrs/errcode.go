package reqerrs

import "github.com/GalvinGao/gofiber-template/pkg/x/urn"

type ErrCode string

// ErrCodeValidationFailed is the error code for validation errors.
const ErrCodeValidationFailed ErrCode = urn.ErrorBase + "validation_failed"

// ErrCodeNotFound is the error code for not found errors.
const ErrCodeNotFound ErrCode = urn.ErrorBase + "not_found"

// ErrCodeInternal is the error code for internal errors.
const ErrCodeInternal ErrCode = urn.ErrorBase + "internal_error_occurred"

// ErrCodeParseFailed is the error code for parsing request errors.
const ErrCodeParseFailed ErrCode = urn.ErrorBase + "parse_failed"

// ErrInvalidRequest is the error code for invalid request errors.
const ErrCodeInvalidRequest ErrCode = urn.ErrorBase + "invalid_request"
