package status

type (
	Status string

	Converter[T any] func(resp T, err error) Status
)

const (
	Success Status = Status(rune(iota))
	ClientError
	ServerError
	Redirect
)

func (s Status) String() string {
	switch s {
	case Success:
		return "success"
	case ClientError:
		return "client_error"
	case ServerError:
		return "server_error"
	case Redirect:
		return "redirect"
	default:
		return "unknown"
	}
}

// DefaultConverter is a default implementation of the Converter interface that returns a status of Success if the error is nil, and ClientError otherwise.
func DefaultConverter[T any](resp T, err error) Status {
	if err != nil {
		return ClientError
	}
	return Success
}

// DefaultHTTPStatusCodeConverter is an implementation of the Converter interface that returns a status based on the HTTP status code.
func DefaultHTTPStatusCodeConverter(statusCode int, err error) Status {
	if statusCode >= 200 && statusCode < 300 {
		return Success
	} else if statusCode >= 300 && statusCode < 400 {
		return Redirect
	} else if statusCode >= 400 && statusCode < 500 {
		return ClientError
	} else {
		return ServerError
	}
}
