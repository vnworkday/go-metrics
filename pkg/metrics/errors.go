package metrics

type ErrTypeConverter func(error) string

func DefaultErrTypeConverter(err error) string {
	return "unknown error"
}
