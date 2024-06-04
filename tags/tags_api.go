package tags

// ****************************************************************************
// This file contains the functions to create tags related to APIs.
// ****************************************************************************

const (
	TagAPIName = "api"
)

// APIName creates a new tag with the key "api" and the given api name as the value.
func APIName(api string) Tag {
	return NewTag(TagAPIName, api)
}
