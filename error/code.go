package error

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
