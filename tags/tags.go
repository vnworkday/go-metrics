package tags

// =============================================================================
// This file contains the functions to create common tags that are not specific to any domain.
// =============================================================================

const (
	TagService   = "service"
	TagOp        = "op"
	TagErrorType = "error_type"
)

// Service creates a new tag with the key "service" and the given service name as the value.
func Service(name string) Tag {
	return NewTag(TagService, name)
}

// Op creates a new tag with the key "op" and the given operation name as the value.
func Op(name string) Tag {
	return NewTag(TagOp, name)
}

// ErrorType creates a new tag with the key "error_type" and the given error type as the value.
func ErrorType(name string) Tag {
	return NewTag(TagErrorType, name)
}
