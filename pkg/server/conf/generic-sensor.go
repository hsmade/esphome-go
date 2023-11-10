package conf

// https://www.home-assistant.io/integrations/sensor/

import (
	"fmt"
	"github.com/hsmade/esphome-go/protobuf"
	"github.com/hsmade/esphome-go/protobuf/api"
	"google.golang.org/protobuf/proto"
	"log/slog"
)

type GenericSensorDefinition struct {
	BaseSensorDefinition
	UnitOfMeasurement string
	AccuracyDecimals  int32
	ForceUpdate       bool
	DeviceClass       string
	StateClass        api.SensorStateClass
	DisabledByDefault bool
	EntityCategory    api.EntityCategory
}

func (B GenericSensorDefinition) ToResponse() ListEntitiesApiResponse {
	return &api.ListEntitiesSensorResponse{
		ObjectId:          B.ObjectId,
		Key:               B.Key,
		Name:              B.Name,
		UniqueId:          B.UniqueId,
		Icon:              B.Icon,
		UnitOfMeasurement: B.UnitOfMeasurement,
		AccuracyDecimals:  B.AccuracyDecimals,
		ForceUpdate:       B.ForceUpdate,
		DeviceClass:       B.DeviceClass,
		StateClass:        B.StateClass,
		DisabledByDefault: B.DisabledByDefault,
		EntityCategory:    B.EntityCategory,
	}
}

func (B GenericSensorDefinition) GetResponseType() protobuf.MsgType {
	return protobuf.ListEntitiesSensorResponseType
}

// GenericSensorState is a message to inform subscribers of updates to states
type GenericSensorState struct {
	BaseSensorState
	State float32
}

func (B GenericSensorState) ToFrame() ([]byte, protobuf.MsgType, error) {
	message := api.SensorStateResponse{
		Key:          B.Key,
		State:        B.State,
		MissingState: B.MissingState,
	}
	slog.Debug("GenericSensorState:ToFrame generating data", "message", fmt.Sprintf("%+v", message))
	data, err := proto.Marshal(&message)
	if err != nil {
		return nil, 0, fmt.Errorf("GenericSensorState:ToFrame: marshalling `SensorStateResponse`: %w", err)
	}

	return data, protobuf.SensorStateResponseType, nil
}
