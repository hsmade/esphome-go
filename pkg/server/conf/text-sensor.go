package conf

// https://www.home-assistant.io/integrations/Text_sensor/

import (
	"fmt"
	"github.com/hsmade/esphome-go/protobuf"
	"github.com/hsmade/esphome-go/protobuf/api"
	"google.golang.org/protobuf/proto"
	"log/slog"
)

type TextSensorDefinition struct {
	BaseSensorDefinition
	DeviceClass        string
	IsStatusTextSensor bool
	DisabledByDefault  bool
	EntityCategory     api.EntityCategory
}

func (B TextSensorDefinition) ToResponse() ListEntitiesApiResponse {
	return &api.ListEntitiesTextSensorResponse{
		ObjectId:          B.ObjectId,
		Key:               B.Key,
		Name:              B.Name,
		UniqueId:          B.UniqueId,
		DisabledByDefault: B.DisabledByDefault,
		Icon:              B.Icon,
		EntityCategory:    B.EntityCategory,
	}
}

func (B TextSensorDefinition) GetResponseType() protobuf.MsgType {
	return protobuf.ListEntitiesTextSensorResponseType
}

// TextSensorState is a message to inform subscribers of updates to states
type TextSensorState struct {
	BaseSensorState
	State string
}

func (B TextSensorState) ToFrame() ([]byte, protobuf.MsgType, error) {
	message := api.TextSensorStateResponse{
		Key:          B.Key,
		State:        B.State,
		MissingState: B.MissingState,
	}
	slog.Debug("TextSensorState:ToFrame generating data", "message", fmt.Sprintf("%+v", message))
	data, err := proto.Marshal(&message)
	if err != nil {
		return nil, 0, fmt.Errorf("TextSensorState:ToFrame: marshalling `TextSensorStateResponse`: %w", err)
	}

	return data, protobuf.TextSensorStateResponseType, nil
}
