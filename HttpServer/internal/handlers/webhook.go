package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WebhookEvent struct {
	EventID string `json:"event_id"`
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

// Validate does basic request validation — Week 4's "request validation"
// topic, kept dependency-free instead of pulling in go-playground/validator.
func (e WebhookEvent) Validate() error {
	if e.EventID == "" {
		return fmt.Errorf("event_id is required")
	}
	if e.Type == "" {
		return fmt.Errorf("type is required")
	}
	return nil
}

// Webhook decodes and validates a JSON body, returning distinct status
// codes for "bad JSON" (400) vs "well-formed but invalid" (422) vs
// success (202 Accepted — the request is acknowledged, not necessarily
// fully processed yet, matching real webhook-ingestion semantics).
func Webhook(w http.ResponseWriter, r *http.Request) {
	var event WebhookEvent

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // reject typos/unexpected fields instead of silently ignoring them
	if err := decoder.Decode(&event); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}

	if err := event.Validate(); err != nil {
		writeError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{
		"status":   "accepted",
		"event_id": event.EventID,
	})
}

func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
