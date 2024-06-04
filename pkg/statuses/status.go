package statuses

type Status string

const (
	Success Status = Status(rune(iota))
	ClientError
	ServerError
	Unknown
)

func (s Status) String() string {
	switch s {
	case Success:
		return "success"
	case ClientError:
		return "client_error"
	case ServerError:
		return "server_error"
	default:
		return "unknown"
	}
}

type StatusConverter[T any] func(resp T, err error) Status

// DefaultConverter is a default implementation of the StatusConverter interface that returns a statuses of Success if the error is nil, and ClientError otherwise.
func DefaultConverter[T any](_ T, err error) Status {
	if err != nil {
		return ClientError
	}
	return Success
}

// DefaultHTTPStatusCodeConverter is an implementation of the StatusConverter interface that returns a statuses based on the HTTP statuses code.
func DefaultHTTPStatusCodeConverter(statusCode int, _ error) Status {
	if statusCode >= 200 && statusCode < 300 {
		return Success
	} else if statusCode >= 400 && statusCode < 500 {
		return ClientError
	} else if statusCode >= 500 && statusCode < 600 {
		return ServerError
	} else {
		return Unknown
	}
}
