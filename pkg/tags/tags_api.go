package tags

// ****************************************************************************
// This file contains the functions to create tags related to APIs.
// ****************************************************************************

const (
	TagAPIName = "api"
	TagAPIOp   = "api_op"
)

// APIName creates a new tag with the key "api" and the given api name as the value.
func APIName(api string) Tag {
	return NewTag(TagAPIName, api)
}

// APIOp creates a new tag with the key "api_op" and the given operation name as the value.
func APIOp(name string) Tag {
	return NewTag(TagAPIOp, name)
}
