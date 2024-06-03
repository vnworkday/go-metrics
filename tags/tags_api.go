package tags

import "github.com/vnworkday/go-metrics/statuses"

// ****************************************************************************
// This file contains the functions to create tags related to APIs.
// ****************************************************************************

const (
	TagAPIName   = "api"
	TagAPIStatus = "api_status"
)

// APIStatus creates a new tag with the key "api_status" and the given statuses as the value.
func APIStatus(status statuses.Status) Tag {
	return NewTag(TagAPIStatus, status.String())
}

// APIName creates a new tag with the key "api" and the given api name as the value.
func APIName(api string) Tag {
	return NewTag(TagAPIName, api)
}
