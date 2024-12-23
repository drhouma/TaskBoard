package route

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

func WriteError(w http.ResponseWriter, errMsg error, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	e := &Error{
		Error: errMsg.Error(),
	}

	body, _ := json.Marshal(e)
	if _, err := w.Write(body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func WriteResponseJson(w http.ResponseWriter, body any) {
	w.Header().Set("Content-Type", "application/json")

	resp, err := json.Marshal(body)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(resp); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
