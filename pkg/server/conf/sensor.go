package conf

import (
	"github.com/hsmade/esphome-go/protobuf"
	"github.com/hsmade/esphome-go/protobuf/api"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Sensor struct {
	Definition SensorDefinition
}

type SensorDefinition interface {
	ToResponse() ListEntitiesApiResponse
	GetResponseType() protobuf.MsgType
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
	ObjectId string
	Key      uint32
	Name     string
	UniqueId string
	Icon     string
}

// ToResponse is used to create an api response struct out of the sensor data
// This implementation is just a stub/example, to satisfy the interface. Sensor structs need to define their own
func (B BaseSensorDefinition) ToResponse() api.ListEntitiesBinarySensorResponse {
	return api.ListEntitiesBinarySensorResponse{}
}

func (B BaseSensorDefinition) GetResponseType() protobuf.MsgType {
	return protobuf.ListEntitiesBinarySensorResponseType
}

type SensorUpdate interface {
	ToFrame() ([]byte, protobuf.MsgType, error)
}

type StateApiResponse interface {
	Reset()
	String() string
	ProtoMessage()
	ProtoReflect() protoreflect.Message
	Descriptor() ([]byte, []int)
	GetKey() uint32
	//GetState() interface{}
	GetMissingState() bool
}

type BaseSensorState struct {
	Key          uint32
	MissingState bool
}

// ToFrame is used to create an api state frame out of the sensor state
// This implementation is just a stub/example, to satisfy the interface. Sensor structs need to define their own
func (S BaseSensorState) ToFrame() ([]byte, protobuf.MsgType, error) {
	return nil, 0, nil
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
