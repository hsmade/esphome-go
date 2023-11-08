package protobuf

import "fmt"

type MsgType uint32

const (
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
	ListEntitiesBinarySensorResponseType      MsgType = 12
	ListEntitiesRequestType                   MsgType = 11
	ListEntitiesDoneResponseType              MsgType = 19
	SubscribeStatesRequestType                MsgType = 20
	SubscribeHomeassistantServicesRequestType MsgType = 34
	SubscribeHomeAssistantStatesRequestType   MsgType = 38

	// 20, 34, 38
)

func (M MsgType) String() string {
	switch M {
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
	case 19:
		return "ListEntitiesDoneResponse"
	case 20:
		return "SubscribeStatesRequest"
	case 34:
		return "SubscribeHomeassistantServicesRequest"
	case 38:
		return "SubscribeHomeAssistantStatesRequest"

	}
	return fmt.Sprintf("unknown:%d", M)
}
