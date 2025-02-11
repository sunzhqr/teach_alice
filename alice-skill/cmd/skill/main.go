package main

import (
	"net/http"

	"github.com/sunzhqr/alice-skill/internal/logger"
	"go.uber.org/zap"
)

func main() {
	parseFlags()
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	if err := logger.Initialize(flagLogLevel); err != nil {
		return err
	}
	logger.Log.Info("Running server", zap.String("address", flagRunAddr))
	// fmt.Println("\"Running\" server on", flagRunAddr)
	return http.ListenAndServe(flagRunAddr, logger.RequestLogger(webhook))
}

func webhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logger.Log.Debug("got request with bad method", zap.String("method", r.Method))
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`
	{
		"response": {
			"text": "Sorry, i don't do anyting"
		},
		"version": "1.0"
	}
	`))
	logger.Log.Debug("Sending HTTP 200 response")
}
