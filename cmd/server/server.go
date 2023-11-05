package main

import (
	"github.com/hsmade/esphome-go/pkg/server"
	"github.com/hsmade/esphome-go/pkg/server/conf"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	var programLevel = new(slog.LevelVar)
	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))
	programLevel.Set(slog.LevelDebug)

	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":9001", nil)

	S := server.Server{
		Port: 6053,
		Config: conf.Config{
			DeviceInfo: conf.DeviceInfo{
				Name: "foobar",
			},
			VerifyPasswordCallback: func(s string) bool {
				return true
			},
		},
	}
	err := S.Listen()
	if err != nil {
		slog.Error("server exited: %w", err)
	}
}
