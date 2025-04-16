package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type statusRecorder struct {
	http.ResponseWriter
	Status int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.Status = code
	rec.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Чтение тела (для повторного использования)
		var bodyBytes []byte
		if r.Body != nil {
			bodyBytes, _ = io.ReadAll(r.Body)
		}
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// Извлекаем CMD и RequestID
		cmd := extractCMDFromRequest(r, bodyBytes)
		requestID := extractRequestIDFromRequest(r, bodyBytes)

		rec := &statusRecorder{ResponseWriter: w, Status: 200}
		next.ServeHTTP(rec, r)

		duration := time.Since(start)
		log.Printf("requestId: %s, cmd: %s, method: %s, path: %s, status: %d, duration: %v",
			requestID, cmd, r.Method, r.URL.Path, rec.Status, duration)
	})
}

func extractCMDFromRequest(r *http.Request, body []byte) string {
	if cmd := r.Header.Get("X-CMD"); cmd != "" {
		return cmd
	}
	if cmd := r.URL.Query().Get("CMD"); cmd != "" {
		return cmd
	}
	var t struct {
		CMD string `json:"CMD"`
	}
	if err := json.Unmarshal(body, &t); err == nil {
		return t.CMD
	}
	return r.Method + " " + r.URL.Path
}

func extractRequestIDFromRequest(r *http.Request, body []byte) string {
	if rid := r.Header.Get("X-Request-ID"); rid != "" {
		return rid
	}
	if rid := r.URL.Query().Get("REQUEST_ID"); rid != "" {
		return rid
	}
	var t struct {
		RequestID string `json:"REQUEST_ID"`
	}
	if err := json.Unmarshal(body, &t); err == nil {
		return t.RequestID
	}
	return "undefined"
}
