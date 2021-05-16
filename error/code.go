package error

import (
	"fmt"
)

// Code represents error code text of Notion
type Code string

const (
	CodeInvalidJson         Code = "invalid_json"
	CodeInvalidRequestURL   Code = "invalid_request_url"
	CodeInvalidRequest      Code = "invalid_request"
	CodeValidationError     Code = "validation_error"
	CodeUnauthorized        Code = "unauthorized"
	CodeRestrictedResource  Code = "restricted_resource"
	CodeObjectNotFound      Code = "object_not_found"
	CodeConflictError       Code = "conflict_error"
	CodeRateLimited         Code = "rate_limited"
	CodeInternalServerError Code = "internal_server_error"
	CodeServiceUnavailable  Code = "service_unavailable"
)

func (c *Code) UnmarshalJSON(data []byte) error {
	code := Code(data[1 : len(data)-1])
	switch code {
	case CodeInvalidJson:
		fallthrough
	case CodeInvalidRequestURL:
		fallthrough
	case CodeInvalidRequest:
		fallthrough
	case CodeValidationError:
		fallthrough
	case CodeUnauthorized:
		fallthrough
	case CodeRestrictedResource:
		fallthrough
	case CodeObjectNotFound:
		fallthrough
	case CodeConflictError:
		fallthrough
	case CodeRateLimited:
		fallthrough
	case CodeInternalServerError:
		fallthrough
	case CodeServiceUnavailable:
		*c = code
		return nil
	}
	return fmt.Errorf("undefined error code: %s", data)
}
