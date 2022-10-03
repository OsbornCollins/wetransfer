// Filename: cmd/api/randomize.go

package main

import (
	"net/http"
	

	"wetransfer.OsbornCollins.net/internal/data"
	"wetransfer.OsbornCollins.net/internal/validator"
)

// createRandomStringHandler for the "GET" /v1/randomize/:int" endpoint
func (app *application) showRandomStringHandler(w http.ResponseWriter, r *http.Request) {
	int1, err := app.readIntParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	userInt := int(int1)
	tools := &data.Tools{
		Int: userInt,
	}

	// initialize a new Validator instance
	v := validator.New()

	//Check the map to determine if there were any validation errors
	if data.ValidateInt(v, tools); !v.Valid(){
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	ranStr := tools.GenerateRandomString(userInt)
	// Adding the int and randomg string to a envelope so we could write it
	// in json format
	jsonText := envelope{
		"int":           userInt,
		"random_string": ranStr,
	}
	err = app.writeJSON(w, http.StatusOK, jsonText, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}
