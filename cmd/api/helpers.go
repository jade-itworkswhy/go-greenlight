package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// setting up all your application-specific handlers and helpers
// so that they are methods on application.
// It helps maintain consistency in your code structure,
// and also future-proofs your code for when those handlers
// and helpers change later and they do need access to a dependency.
func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	// try encoding data to json
	// returns []byte slice(containing encoded JSON)
	// alphabetically sorted
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	js = append(js, '\n')

	// construct additional header
	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
