package api

import (
	"encoding/json"
	"net/http"

	log "github.com/elleFlorio/gru/Godeps/_workspace/src/github.com/Sirupsen/logrus"

	cfg "github.com/elleFlorio/gru/configuration"
)

// /gru/v1/node
func GetInfoNode(w http.ResponseWriter, r *http.Request) {
	info := cfg.GetNode()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(info); err != nil {
		log.WithFields(log.Fields{
			"status":  "http response",
			"request": "GetInfoNode",
			"error":   err,
		}).Errorln("API Server")
	}
}
