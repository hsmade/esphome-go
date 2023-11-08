package conf

import (
	"github.com/hsmade/esphome-go/protobuf"
	"github.com/hsmade/esphome-go/protobuf/api"
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

type BinarySensorState struct {
	BaseSensorState
	State bool
}

func (B BinarySensorState) ToResponse() StateApiResponse {
	return &api.BinarySensorStateResponse{
		Key:          B.Key,
		State:        B.State,
		MissingState: B.MissingState,
	}
}

func (B BinarySensorState) GetResponseType() protobuf.MsgType {
	return protobuf.ListEntitiesBinarySensorResponseType
}
