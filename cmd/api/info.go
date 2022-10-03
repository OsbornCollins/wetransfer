// Filename: cmd/api/info.go

package main

import (
	"fmt"
	"net/http"

	"wetransfer.OsbornCollins.net/internal/data"
	"wetransfer.OsbornCollins.net/internal/validator"
)

// createInfoHandler for the "POST" /v1/info" endpoint
func (app *application) createInfoHandler(w http.ResponseWriter, r *http.Request) {
	// Our Target decode destination
	var input struct {
		Name       string `json:"name"`
		Occupation string `json:"occupation,omitempty"`
		Phone      string `json:"phone"`
		Email      string `json:"email,omitempty"`
		Country    string `json:"country"`
		Address    string `json:"address"`
	}
	// Initialize a new json.Decoder instance
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	//Copy the values from the input struct to a new info struct
	info := &data.Info{
		Name:       input.Name,
		Occupation: input.Occupation,
		Phone:      input.Phone,
		Email:      input.Email,
		Country:    input.Country,
		Address:    input.Address,
	}
	// initialize a new Validator instance
	v := validator.New()

	//Check the map to determine if there were any validation errors
	if data.ValidateInfo(v, info); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Display the request
	fmt.Fprintf(w, "%+v\n", input)
}
