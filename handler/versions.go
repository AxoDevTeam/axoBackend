package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AxoDevTeam/axoBackend/config"
)

var Versions = http.HandlerFunc(versions)

func versions(w http.ResponseWriter, r *http.Request) {
	vers, ok := config.Conf["ver"].(map[string]interface{})
	if !ok {
		http.Error(w, "type fail", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(vers); err != nil {
		http.Error(w, "encode fail", http.StatusInternalServerError)
	}
}
