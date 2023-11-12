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
	ListEntitiesSensorResponseType            MsgType = 16
	ListEntitiesTextSensorResponseType        MsgType = 18
	ListEntitiesDoneResponseType              MsgType = 19
	SubscribeStatesRequestType                MsgType = 20
	BinarySensorStateResponseType             MsgType = 21
	SensorStateResponseType                   MsgType = 25
	TextSensorStateResponseType               MsgType = 27
	SubscribeHomeassistantServicesRequestType MsgType = 34
	HomeassistantServiceResponseType          MsgType = 35
	SubscribeHomeAssistantStatesRequestType   MsgType = 38
	SubscribeHomeAssistantStateResponseType   MsgType = 39
	HomeAssistantStateResponseType            MsgType = 40
	ListEntitiesServicesResponseType          MsgType = 41
	ExecuteServiceRequestType                 MsgType = 42
	ListEntitiesCameraResponseType            MsgType = 43
	CameraImageResponseType                   MsgType = 44
	CameraImageRequestType                    MsgType = 45
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
	case 16:
		return "ListEntitiesSensorResponse"
	case 18:
		return "ListEntitiesTextSensorResponse"
	case 19:
		return "ListEntitiesDoneResponse"
	case 20:
		return "SubscribeStatesRequest"
	case 21:
		return "BinarySensorStateResponse"
	case 25:
		return "SensorStateResponse"
	case 27:
		return "TextSensorStateResponse"
	case 34:
		return "SubscribeHomeassistantServicesRequest"
	case 35:
		return "HomeassistantServiceResponse"
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

	}
	return fmt.Sprintf("unknown:%d", M)
}
