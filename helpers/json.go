package helpers

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func JSON(key, value string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp[key] = value
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		slog.Error("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}
