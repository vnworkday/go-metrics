package tags

import "github.com/vnworkday/go-metrics/pkg/statuses"

// =============================================================================
// This file contains the functions to create common tags that are not specific to any domain.
// =============================================================================

const (
	TagService   = "service"
	TagStatus    = "status"
	TagErrorType = "error_type"
)

// Service creates a new tag with the key "service" and the given service name as the value.
func Service(name string) Tag {
	return NewTag(TagService, name)
}

// Status creates a new tag with the key "api_status" and the given statuses as the value.
func Status(status statuses.Status) Tag {
	return NewTag(TagStatus, status.String())
}

// ErrorType creates a new tag with the key "error_type" and the given error type as the value.
func ErrorType(name string) Tag {
	return NewTag(TagErrorType, name)
}
