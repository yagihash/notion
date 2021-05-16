package error

type Code string

const (
	InvalidJson         Code = "invalid_json"
	InvalidRequestURL   Code = "invalid_request_url"
	InvalidRequest      Code = "invalid_request"
	ValidationError     Code = "validation_error"
	Unauthorized        Code = "unauthorized"
	RestrictedResource  Code = "restricted_resource"
	ObjectNotFound      Code = "object_not_found"
	ConflictError       Code = "conflict_error"
	InternalServerError Code = "internal_server_error"
	ServiceUnavailable  Code = "service_unavailable"
)

var ()
