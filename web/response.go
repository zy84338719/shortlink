package web

import (
	"encoding/json"
	"io"
	"net/http"
	"shortlink/models"
)

func sendErrorResponse(w http.ResponseWriter, errReps models.Response) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(errReps.Status)
	marshal, _ := json.Marshal(errReps.Result)
	_, _ = io.WriteString(w, string(marshal))
}

func sendNormalResponse(w http.ResponseWriter, result models.Result, status int) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	marshal, _ := json.Marshal(result)
	_, _ = io.WriteString(w, string(marshal))
}
