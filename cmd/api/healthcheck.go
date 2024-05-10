package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	/* fixed-string approach

	// string literal
	js := `{"status": "available", "environment": %q, "version": %q}`
	js = fmt.Sprintf(js, app.config.env, version)

	// as json
	w.Header().Set("Content-Type", "application/json")

	// write to the body
	w.Write([]byte(js))
	*/

	// json approach

	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"status":      "available",
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
