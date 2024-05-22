package data

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Runtime int32

func (r *Runtime) MarshalJSON(jsonValue []byte) error {
	// unquote the string from the input jsonValue param
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))

	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	// split the runtime string to isolate number
	parts := strings.Split(unquotedJSONValue, " ")

	// sanity check(expected example: 3 mins)
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)

	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(i)
	return nil
}
