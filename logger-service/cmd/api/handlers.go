package main

import (
	"logger-service/data"
	"net/http"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	// read json into var

	var requestPayload JSONPayload
	_ = app.readJSON(w, r, &requestPayload)

	// insert Data

	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}

	response := jsonResponse{
		Error:   false,
		Message: "Log entry inserted successfully",
	}

	app.writeJSON(w, http.StatusAccepted, response)
}
