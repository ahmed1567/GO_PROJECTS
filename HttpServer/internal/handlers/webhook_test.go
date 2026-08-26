package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// A table-driven test: one test FUNCTION, a table (slice) of CASES, and a
// loop that runs the same assertions against every case via t.Run — Go's
// idiomatic alternative to writing a separate test function per scenario.
func TestWebhook(t *testing.T) {
	tests := []struct {
		name       string
		body       string
		wantStatus int
	}{
		{
			name:       "valid event",
			body:       `{"event_id":"evt_1","type":"payment.success","payload":"data"}`,
			wantStatus: http.StatusAccepted,
		},
		{
			name:       "missing event_id",
			body:       `{"type":"payment.success"}`,
			wantStatus: http.StatusUnprocessableEntity,
		},
		{
			name:       "missing type",
			body:       `{"event_id":"evt_1"}`,
			wantStatus: http.StatusUnprocessableEntity,
		},
		{
			name:       "invalid JSON",
			body:       `{not valid json`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "unknown field rejected",
			body:       `{"event_id":"evt_1","type":"x","extra_field":"nope"}`,
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/webhook", strings.NewReader(tt.body))
			rec := httptest.NewRecorder()

			Webhook(rec, req)

			if rec.Code != tt.wantStatus {
				t.Errorf("got status %d, want %d (body: %s)", rec.Code, tt.wantStatus, rec.Body.String())
			}
		})
	}
}

func TestHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	Health(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("got status %d, want %d", rec.Code, http.StatusOK)
	}
}
