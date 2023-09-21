package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Message string   `json:"message"`
	Details []string `json:"details,omitempty`
}

func RespondJSON(ctx context.Context, w http.ResponseWriter, body any, status int) error {
	w.Header().Set("Context-Type", "application/json")
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := ErrorResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			return fmt.Errorf("RespondJSON failed: %w", err)
		}
		return err
	}

	w.WriteHeader(status)
	if _, err := fmt.Fprintf(w, "%s", bodyBytes); err != nil {
		return fmt.Errorf("RespondJSON failed: %w", err)
	}
	return nil
}
