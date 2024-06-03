package warnings

// ****************************************************************************
// This file contains the functions to create warnings related to tags.
// ****************************************************************************

const (
	WarningTagCollectionLimitReachedMessage = "CollectionLimitReachedMessage"
	WarningTagNameSizeLimitReachedMessage   = "NameSizeLimitReachedMessage"
	WarningTagValueSizeLimitReachedMessage  = "ValueSizeLimitReachedMessage"
)

// TagOverflow used when the number of tags exceeds the limit.
func TagOverflow(metric, tagName, tagValue string) Warning {
	return NewWarning(metric, WarningTagCollectionLimitReachedMessage).addTag(tagName, tagValue)
}

// TagNameTooLong used when the tag name exceeds the length limit.
func TagNameTooLong(metric, tagName, tagValue string) Warning {
	return NewWarning(metric, WarningTagNameSizeLimitReachedMessage).addTag(tagName, tagValue)
}

// TagValueTooLong used when the tag value exceeds the length limit.
func TagValueTooLong(metric, tagName, tagValue string) Warning {
	return NewWarning(metric, WarningTagValueSizeLimitReachedMessage).addTag(tagName, tagValue)
}
