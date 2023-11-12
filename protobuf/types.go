package protobuf

import "fmt"

type MsgType uint32

const (
	InvalidType                               MsgType = 0
	HelloRequestType                          MsgType = 1
	HelloResponseType                         MsgType = 2
	ConnectRequestType                        MsgType = 3
	ConnectResponseType                       MsgType = 4
	DisconnectRequestType                     MsgType = 5
	DisconnectResponseType                    MsgType = 6
	PingRequestType                           MsgType = 7
	PingResponseType                          MsgType = 8
	DeviceInfoRequestType                     MsgType = 9
	DeviceInfoResponseType                    MsgType = 10
	ListEntitiesRequestType                   MsgType = 11
	ListEntitiesBinarySensorResponseType      MsgType = 12
	ListEntitiesCoverResponseType             MsgType = 13
	ListEntitiesFanResponseType               MsgType = 14
	ListEntitiesLightResponseType             MsgType = 15
	ListEntitiesSensorResponseType            MsgType = 16
	ListEntitiesSwitchResponseType            MsgType = 17
	ListEntitiesTextSensorResponseType        MsgType = 18
	ListEntitiesDoneResponseType              MsgType = 19
	SubscribeStatesRequestType                MsgType = 20
	BinarySensorStateResponseType             MsgType = 21
	CoverStateResponseType                    MsgType = 22
	FanStateResponseType                      MsgType = 23
	LightStateResponseType                    MsgType = 24
	SensorStateResponseType                   MsgType = 25
	SwitchStateResponseType                   MsgType = 26
	TextSensorStateResponseType               MsgType = 27
	SubscribeLogsRequestType                  MsgType = 28
	SubscribeLogsResponseType                 MsgType = 29
	CoverCommandRequestType                   MsgType = 30
	FanCommandRequestType                     MsgType = 31
	LightCommandRequestType                   MsgType = 32
	SwitchCommandRequestType                  MsgType = 33
	SubscribeHomeassistantServicesRequestType MsgType = 34
	HomeassistantServiceResponseType          MsgType = 35
	GetTimeRequestType                        MsgType = 36
	GetTimeResponseType                       MsgType = 37
	SubscribeHomeAssistantStatesRequestType   MsgType = 38
	SubscribeHomeAssistantStateResponseType   MsgType = 39
	HomeAssistantStateResponseType            MsgType = 40
	ListEntitiesServicesResponseType          MsgType = 41
	ExecuteServiceRequestType                 MsgType = 42
	ListEntitiesCameraResponseType            MsgType = 43
	CameraImageResponseType                   MsgType = 44
	CameraImageRequestType                    MsgType = 45
	ListEntitiesClimateResponseType           MsgType = 46
	ClimateStateResponseType                  MsgType = 47
	ClimateCommandRequestType                 MsgType = 48
	ListEntitiesNumberResponseType            MsgType = 49
	NumberStateResponseType                   MsgType = 50
	NumberCommandRequestType                  MsgType = 51
	ListEntitiesSelectResponseType            MsgType = 52
	SelectStateResponseType                   MsgType = 53
	SelectCommandRequestType                  MsgType = 54
)

func (M MsgType) String() string {
	switch M {
	case 0:
		return "Invalid"
	case 1:
		return "HelloRequest"
	case 2:
		return "HelloResponse"
	case 3:
		return "ConnectRequest"
	case 4:
		return "ConnectResponse"
	case 5:
		return "DisconnectRequest"
	case 6:
		return "DisconnectResponse"
	case 7:
		return "PingRequest"
	case 8:
		return "PingResponse"
	case 9:
		return "DeviceInfoRequest"
	case 10:
		return "DeviceInfoResponse"
	case 11:
		return "ListEntitiesRequest"
	case 12:
		return "ListEntitiesBinarySensorResponse"
	case 13:
		return "ListEntitiesCoverResponse"
	case 14:
		return "ListEntitiesFanResponse"
	case 15:
		return "ListEntitiesLightResponse"
	case 16:
		return "ListEntitiesSensorResponse"
	case 17:
		return "ListEntitiesSwitchResponse"
	case 18:
		return "ListEntitiesTextSensorResponse"
	case 19:
		return "ListEntitiesDoneResponse"
	case 20:
		return "SubscribeStatesRequest"
	case 21:
		return "BinarySensorStateResponse"
	case 22:
		return "CoverStateResponse"
	case 23:
		return "FanStateResponse"
	case 24:
		return "LightStateResponse"
	case 25:
		return "SensorStateResponse"
	case 26:
		return "SwitchStateResponse"
	case 27:
		return "TextSensorStateResponse"
	case 28:
		return "SubscribeLogsRequest"
	case 29:
		return "SubscribeLogsResponse"
	case 30:
		return "CoverCommandRequest"
	case 31:
		return "FanCommandRequest"
	case 32:
		return "LightCommandRequest"
	case 33:
		return "SwitchCommandRequest"
	case 34:
		return "SubscribeHomeassistantServicesRequest"
	case 35:
		return "HomeassistantServiceResponse"
	case 36:
		return "GetTimeRequest"
	case 37:
		return "GetTimeResponse"
	case 38:
		return "SubscribeHomeAssistantStatesRequest"
	case 39:
		return "SubscribeHomeAssistantStateResponse"
	case 40:
		return "HomeAssistantStateResponse"
	case 41:
		return "ListEntitiesServicesResponse"
	case 42:
		return "ExecuteServiceRequest"
	case 43:
		return "ListEntitiesCameraResponse"
	case 44:
		return "CameraImageResponse"
	case 45:
		return "CameraImageRequest"
	case 46:
		return "ListEntitiesClimateResponse"
	case 47:
		return "ClimateStateResponse"
	case 48:
		return "ClimateCommandRequest"
	case 49:
		return "ListEntitiesNumberResponse"
	case 50:
		return "NumberStateResponse"
	case 51:
		return "NumberCommandRequest"
	case 52:
		return "ListEntitiesSelectResponse"
	case 53:
		return "SelectStateResponse"
	case 54:
		return "SelectCommandRequest"
	}
	return fmt.Sprintf("unknown:%d", M)
}
