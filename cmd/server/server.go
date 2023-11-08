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
	binarySensorExampleUpdates := make(chan conf.SensorUpdate, 1) // 1 to not block when we send to this below
	binarySensorExample := conf.Sensor{
		Definition: conf.BinarySensorDefinition{
			BaseSensorDefinition: conf.BaseSensorDefinition{
				ObjectId: "test",
				Key:      1,
				Name:     "test",
				UniqueId: "test",
			},
		},
		Updates: binarySensorExampleUpdates,
	}
	binarySensorExampleUpdates <- conf.BinarySensorState{
		BaseSensorState: conf.BaseSensorState{
			Key:          1,
			MissingState: false,
		},
		State: true,
	}

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
				binarySensorExample,
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
