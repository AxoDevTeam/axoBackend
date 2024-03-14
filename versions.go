package main

import (
	"encoding/json"
	"net/http"
)

func versions(w http.ResponseWriter, r *http.Request) {
	vers, ok := config["ver"].(map[string]interface{})
	if !ok {
		http.Error(w, "type fail", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(vers); err != nil {
		http.Error(w, "encode fail", http.StatusInternalServerError)
	}
}
