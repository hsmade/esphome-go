package main

import (
	"github.com/hsmade/esphome-go/pkg/server"
	"github.com/hsmade/esphome-go/pkg/server/conf"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log/slog"
	"math/rand"
	"net/http"
	"os"
	"time"
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
	binarySensorExample := conf.Sensor{
		Definition: conf.BinarySensorDefinition{
			BaseSensorDefinition: conf.BaseSensorDefinition{
				ObjectId: "test",
				Key:      1,
				Name:     "test",
				UniqueId: "test",
			},
		},
	}

	// define server object and config
	binarySensorExampleUpdates := make(chan conf.SensorUpdate, 1)
	S := server.Server{
		Port: 6053,
		Config: conf.Config{
			DeviceInfo: conf.DeviceInfo{
				Name:         "foobar",
				UsesPassword: true,
			},
			Updates: binarySensorExampleUpdates,
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
			slog.Error("main: server exited: %w", err)
		}
	}()

	time.Sleep(5 * time.Second)
	// send an update to a sensor
	tick := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-tick.C:
			slog.Info("main: sending update for binary sensor")
			binarySensorExampleUpdates <- conf.BinarySensorState{
				BaseSensorState: conf.BaseSensorState{
					Key:          uint32(rand.Intn(1)),
					MissingState: false,
				},
				State: true,
			}
		}
	}
}
