package conf

import (
	"github.com/hsmade/esphome-go/protobuf/api"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// FIXME: how to expose / ingest sensors and their values?

type Sensor struct {
	Definition SensorDefinition
	Updates    chan SensorUpdate
	Commands   chan SensorCommand
}

type SensorDefinition interface {
	toResponse() ListEntitiesApiResponse
}

// ListEntitiesApiResponse describes the minimal interface for ListEntitiesXxxxxSensorResponse objects
type ListEntitiesApiResponse interface {
	Reset()
	String() string
	ProtoMessage()
	ProtoReflect() protoreflect.Message
	Descriptor() ([]byte, []int)
	GetObjectId() string
	GetKey() uint32
	GetName() string
	GetUniqueId() string
	GetIcon() string
}

type BaseSensorDefinition struct {
	ObjectId string `json:"object_id,omitempty"`
	Key      uint32 `json:"key,omitempty"`
	Name     string `json:"name,omitempty"`
	UniqueId string `json:"unique_id,omitempty"`
	Icon     string `json:"icon,omitempty"`
}

// toResponse is used to create an api response struct out of the sensor data
// This implementation is just a stub/example, to satisfy the interface. Sensor structs need to define their own
func (B BaseSensorDefinition) toResponse() api.ListEntitiesBinarySensorResponse {
	return api.ListEntitiesBinarySensorResponse{}
}

type SensorUpdate interface {
	toResponse() StateApiResponse
}

type StateApiResponse interface {
	Reset()
	String() string
	ProtoMessage()
	ProtoReflect() protoreflect.Message
	Descriptor() ([]byte, []int)
	GetKey() uint32
	GetState() interface{}
	GetMissingState() bool
}

type BaseSensorState struct {
	Key          uint32
	MissingState bool
}

// toResponse is used to create an api state response struct out of the sensor state
// This implementation is just a stub/example, to satisfy the interface. Sensor structs need to define their own
func (S BaseSensorState) toResponse() api.BinarySensorStateResponse {
	return api.BinarySensorStateResponse{}
}

type SensorCommand interface {
	toCommand() CommandApiCommand
}

type CommandApiCommand interface {
	Reset()
	String() string
	ProtoMessage()
	ProtoReflect() protoreflect.Message
	Descriptor() ([]byte, []int)
	GetKey() uint32
	GetHasState() bool
	GetState() interface{}
}

type BaseSensorCommand struct {
	Key      uint32
	HasState bool
	State    bool
}

// toCommand is used to create an api command request struct out of the sensor command
// This implementation is just a stub/example, to satisfy the interface. Sensor structs need to define their own
func (S BaseSensorState) toCommand() api.LightCommandRequest {
	return api.LightCommandRequest{}
}
