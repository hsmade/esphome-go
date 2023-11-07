package conf

import "github.com/hsmade/esphome-go/protobuf/api"

// FIXME: how to expose / ingest sensors and their values?

type Sensor struct {
	Definition SensorDefinition
	//Updates    chan SensorUpdate
	//Commands   chan SensorCommand
}

type SensorDefinition interface {
	bla() // FIXME
}

type BaseSensor struct {
	ObjectId string `json:"object_id,omitempty"`
	Key      uint32 `json:"key,omitempty"`
	Name     string `json:"name,omitempty"`
	UniqueId string `json:"unique_id,omitempty"`
	Icon     string `json:"icon,omitempty"`
}

func (B BaseSensor) bla() {}

type BinarySensor struct {
	BaseSensor
	DeviceClass          string             `json:"device_class,omitempty"`
	IsStatusBinarySensor bool               `json:"is_status_binary_sensor,omitempty"`
	DisabledByDefault    bool               `json:"disabled_by_default,omitempty"`
	Icon                 string             `json:"icon,omitempty"`
	EntityCategory       api.EntityCategory `json:"entity_category,omitempty"`
}

func test() {
	_ = Sensor{
		Definition: BinarySensor{},
	}
}
