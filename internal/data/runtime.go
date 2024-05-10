package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) { // will work for both value and pointer argument
	// custom string for runtime formatting
	jsonValue := fmt.Sprintf("%d mins", r)

	// wrap with double quotes for valid json
	quotedJSONValue := strconv.Quote(jsonValue)

	// return byte slice
	return []byte(quotedJSONValue), nil
}
