package handlers

import (
	"encoding/json"
	"net/http"
)

// Health is a liveness check — intentionally has NO auth middleware
// applied to it in cmd/server/main.go, since load balancers/orchestrators
// need to hit it without a key.
func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
