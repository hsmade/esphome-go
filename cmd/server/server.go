package main

import (
	"github.com/hsmade/esphome-go/pkg/server"
	"github.com/hsmade/esphome-go/pkg/server/conf"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log/slog"
	"math/rand"
	"net/http"
	"os"
)

func main() {
	// set up logging (optional)
	var programLevel = new(slog.LevelVar)
	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))
	programLevel.Set(slog.LevelDebug)

	// set up metrics (optional)
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":9001", nil)

	// define a binary sensor
	//binarySensorExample := api.ListEntitiesBinarySensorResponse{}

	// define server object and config
	S := server.Server{
		Port: 6053,
		Config: conf.Config{
			DeviceInfo: conf.DeviceInfo{
				Name:         "foobar",
				UsesPassword: true,
			},
			VerifyPasswordCallback: func(s string) bool {
				// password checking logic goes here
				return true
			},
			Sensors: []conf.Sensor{
				conf.Sensor{
					Definition: conf.BinarySensorDefinition{},
				},
			},
		},
	}

	// and finally start the server
	go func() {
		err := S.Listen()
		if err != nil {
			slog.Error("server exited: %w", err)
		}
	}()

	for {
		_ = rand.Intn(1) == 1 // randomly set the binary sensor's value
	}
}
