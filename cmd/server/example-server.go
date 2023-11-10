package main

import (
	"github.com/hsmade/esphome-go/pkg/server"
	"github.com/hsmade/esphome-go/pkg/server/conf"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"hash/fnv"
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

	// define a binary sensor https://www.home-assistant.io/integrations/binary_sensor/
	// we need a key for our sensor
	key := fnv.New32()
	_, _ = key.Write([]byte("test_binary")) // name of sensor, should be camelcase
	binaryKey := key.Sum32()

	binarySensorExample := conf.Sensor{
		Definition: conf.BinarySensorDefinition{
			BaseSensorDefinition: conf.BaseSensorDefinition{
				ObjectId: "esphome_go_test_binary",
				Key:      binaryKey,
				Name:     "test_binary",
				UniqueId: "esphome_go_test_binary",
			},
		},
	}

	// define a generic sensor https://www.home-assistant.io/integrations/sensor/
	// we need a key for our sensor
	key = fnv.New32()
	_, _ = key.Write([]byte("test_sensor")) // name of sensor, should be camelcase
	genericKey := key.Sum32()

	sensorExample := conf.Sensor{
		Definition: conf.GenericSensorDefinition{
			BaseSensorDefinition: conf.BaseSensorDefinition{
				ObjectId: "esphome_go_test_sensor",
				Key:      genericKey,
				Name:     "test_sensor",
				UniqueId: "esphome_go_test_sensor",
			},
			DisabledByDefault: false,
			UnitOfMeasurement: "V",
			AccuracyDecimals:  2,
			DeviceClass:       "voltage", // https://www.home-assistant.io/docs/configuration/customizing-devices/#device-class
		},
	}

	// define server object and config
	SensorExampleUpdates := make(chan conf.SensorUpdate, 1)
	S := server.Server{
		Port: 6053,
		Config: conf.Config{
			DeviceInfo: conf.DeviceInfo{
				Name:         "esphome-go test",
				UsesPassword: true,
			},
			Updates: SensorExampleUpdates,
			VerifyPasswordCallback: func(s string) bool {
				// password checking logic goes here
				return true
			},
			Sensors: []conf.Sensor{
				binarySensorExample,
				sensorExample,
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

	// send an update to a sensor every 5 seconds
	// Normally you should only send an update, if the state is actually updated :)
	tick := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-tick.C:
			slog.Info("main: sending update for binary sensor")
			SensorExampleUpdates <- conf.BinarySensorState{
				BaseSensorState: conf.BaseSensorState{
					Key:          binaryKey,
					MissingState: false,
				},
				State: rand.Intn(2) == 1, // randomly on/off
			}
			slog.Info("main: sending update for generic sensor")
			SensorExampleUpdates <- conf.GenericSensorState{
				BaseSensorState: conf.BaseSensorState{
					Key:          genericKey,
					MissingState: false,
				},
				State: rand.Float32(),
			}
		}
	}
}
