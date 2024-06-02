package warnings

import (
	"fmt"
	"log"
	"strings"
)

const (
	WarnCardinalityLimitReachedMessage = "CardinalityLimitReachedMessage"
	WarnNameSizeLimitReachedMessage    = "NameSizeLimitReachedMessage"
	WarnValueSizeLimitReachedMessage   = "ValueSizeLimitReachedMessage"
)

type (
	// Warning represents a warning message that is generated during the processing of a metric.
	Warning struct {
		// The name of the metric
		metric string
		// The warning message
		message string
		// The tags of the metric
		tags map[string]string
	}

	// Handler used to handle the warnings generated during the processing of a metric.
	Handler func(warning Warning)
)

func DefaultWarningHandler() Handler {
	return func(warning Warning) {
		log.Default().Println(warning.String())
	}
}

func TagOverflow(metric, tagName, tagValue string) Warning {
	return newWarning(metric, WarnCardinalityLimitReachedMessage).withTag(tagName, tagValue)
}

func TagNameTooLong(metric, tagName, tagValue string) Warning {
	return newWarning(metric, WarnNameSizeLimitReachedMessage).withTag(tagName, tagValue)
}

func TagValueTooLong(metric, tagName, tagValue string) Warning {
	return newWarning(metric, WarnValueSizeLimitReachedMessage).withTag(tagName, tagValue)
}

func InvalidUnit(metric, unit string) Warning {
	return newWarning(metric, fmt.Sprintf("invalid unit \"%s\" will be ignored", unit))
}

func newWarning(metric, message string) Warning {
	return Warning{
		metric:  metric,
		message: message,
		tags:    make(map[string]string),
	}
}

func (w Warning) withTag(name, value string) Warning {
	w.tags[name] = value
	return w
}

func (w Warning) Metric() string {
	return w.metric
}

func (w Warning) Message() string {
	return w.message
}

func (w Warning) Tags() map[string]string {
	return w.tags
}

// String returns a string representation of the warning in the format:
// metric="name" message="message" label_name="name" label_value="value".
func (w Warning) String() string {
	var sb strings.Builder
	sb.WriteString(printout("metric", w.metric))
	sb.WriteString(" ")
	sb.WriteString(printout("message", w.message))
	for k, v := range w.tags {
		sb.WriteString(" ")
		sb.WriteString(printout("label_name", k))
		sb.WriteString(" ")
		sb.WriteString(printout("label_value", v))
	}
	return sb.String()
}

// printout returns a string representation of a tag in the format key="value".
func printout(key, val string) string {
	return fmt.Sprintf(`%s="%s"`, key, val)
}
