package conf

import (
	"fmt"
	"github.com/hsmade/esphome-go/protobuf"
	"github.com/hsmade/esphome-go/protobuf/api"
	"google.golang.org/protobuf/proto"
	"log/slog"
)

type BinarySensorDefinition struct {
	BaseSensorDefinition
	DeviceClass          string             `json:"device_class,omitempty"`
	IsStatusBinarySensor bool               `json:"is_status_binary_sensor,omitempty"`
	DisabledByDefault    bool               `json:"disabled_by_default,omitempty"`
	Icon                 string             `json:"icon,omitempty"`
	EntityCategory       api.EntityCategory `json:"entity_category,omitempty"`
}

func (B BinarySensorDefinition) ToResponse() ListEntitiesApiResponse {
	return &api.ListEntitiesBinarySensorResponse{
		ObjectId:             B.ObjectId,
		Key:                  B.Key,
		Name:                 B.Name,
		UniqueId:             B.UniqueId,
		DeviceClass:          B.DeviceClass,
		IsStatusBinarySensor: B.IsStatusBinarySensor,
		DisabledByDefault:    B.DisabledByDefault,
		Icon:                 B.Icon,
		EntityCategory:       B.EntityCategory,
	}
}

// BinarySensorState is a message to inform subscribers of updates to states
type BinarySensorState struct {
	BaseSensorState
	State bool
}

func (B BinarySensorState) ToFrame() ([]byte, protobuf.MsgType, error) {
	message := api.BinarySensorStateResponse{
		Key:          B.Key,
		State:        B.State,
		MissingState: B.MissingState,
	}
	slog.Debug("BinarySensorState:ToFrame generating data", "message", fmt.Sprintf("%+v", message))
	data, err := proto.Marshal(&message)
	if err != nil {
		return nil, 0, fmt.Errorf("BinarySensorState:ToFrame: marshalling `BinarySensorStateResponse`: %w", err)
	}

	return data, protobuf.BinarySensorStateResponseType, nil
}

func (B BinarySensorState) GetResponseType() protobuf.MsgType {
	return protobuf.ListEntitiesBinarySensorResponseType
}
