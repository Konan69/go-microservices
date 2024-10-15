package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
	Log    LogPayload  `json:"log,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {

	payload := jsonResponse{
		Error:   false,
		Message: "Broker is up and running",
		Data:    nil,
	}

	_ = app.writeJSON(w, http.StatusOK, payload)

}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {

	var RequestPayload RequestPayload

	err := app.readJSON(w, r, &RequestPayload)
	if err != nil {
		app.errorJson(w, err)
		return
	}
	switch RequestPayload.Action {
	case "auth":
		app.authenticate(w, &RequestPayload.Auth)
		return
	case "log":
		app.logItem(w, &RequestPayload.Log)
	default:
		app.errorJson(w, errors.New("unknown action"))
		return
	}
}

func (app *Config) authenticate(w http.ResponseWriter, a *AuthPayload) {

	jsonData, _ := json.MarshalIndent(a, "", "\t")

	request, err := http.NewRequest("POST", "http://auth-service/authenticate", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJson(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJson(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJson(w, errors.New("error calling auth service"))
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)

	if err != nil {
		app.errorJson(w, errors.New("error decoding response from auth service"))
		return
	}

	if jsonFromService.Error {
		app.errorJson(w, errors.New(jsonFromService.Message))
		return
	}

	var payload jsonResponse

	payload.Error = false
	payload.Message = "Authentication successful"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusOK, payload)

}

func (app *Config) logItem(w http.ResponseWriter, l *LogPayload) {
	jsonData, _ := json.MarshalIndent(l, "", "\t")
	request, err := http.NewRequest("POST", "http://logger-service/log", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJson(w, err)
		return
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJson(w, err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		app.errorJson(w, err)
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Log entry inserted successfully"

	app.writeJSON(w, http.StatusAccepted, payload)
}
