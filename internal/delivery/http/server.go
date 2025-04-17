package http

import (
	"azyk/util/logger"
	"net/http"
)

// StartServer запускает HTTP сервер
func StartServer(router *Router, addr string) {
	logger.Log.Infof("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		logger.Log.WithField("error", err.Error()).Fatal("Failed to start server")
	}
}
