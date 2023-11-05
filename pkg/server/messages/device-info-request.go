package messages

import "C"
import (
	"fmt"
	"github.com/hsmade/esphome-go/pkg/server/conf"
	"github.com/hsmade/esphome-go/pkg/server/frames"
	"github.com/hsmade/esphome-go/protobuf"
	"github.com/hsmade/esphome-go/protobuf/api"
	"google.golang.org/protobuf/proto"
	"net"
)

type DeviceInfo struct {
	err     error
	payload *api.DeviceInfoRequest
}

func DeviceInfoRequest(frame *frames.Frame) DeviceInfo {
	// parse frame into deviceInfo object
	h := DeviceInfo{
		payload: &api.DeviceInfoRequest{},
	}
	h.err = proto.Unmarshal(frame.Data, h.payload)
	return h
}

func (H DeviceInfo) Respond(conn net.Conn, config conf.Config) error {
	if H.err != nil {
		return fmt.Errorf("failed parsing DeviceInfoRequest frame: %w", H.err)
	}

	response := api.DeviceInfoResponse{
		UsesPassword:                config.DeviceInfo.UsesPassword,
		Name:                        config.DeviceInfo.Name,
		MacAddress:                  config.DeviceInfo.MacAddress,
		EsphomeVersion:              config.DeviceInfo.EsphomeVersion,
		CompilationTime:             config.DeviceInfo.CompilationTime,
		Model:                       config.DeviceInfo.Model,
		HasDeepSleep:                config.DeviceInfo.HasDeepSleep,
		ProjectName:                 config.DeviceInfo.ProjectName,
		ProjectVersion:              config.DeviceInfo.ProjectVersion,
		WebserverPort:               config.DeviceInfo.WebserverPort,
		LegacyBluetoothProxyVersion: config.DeviceInfo.LegacyBluetoothProxyVersion,
		BluetoothProxyFeatureFlags:  config.DeviceInfo.BluetoothProxyFeatureFlags,
		Manufacturer:                config.DeviceInfo.Manufacturer,
		FriendlyName:                config.DeviceInfo.FriendlyName,
		VoiceAssistantVersion:       config.DeviceInfo.VoiceAssistantVersion,
		SuggestedArea:               config.DeviceInfo.SuggestedArea,
	}
	data, err := proto.Marshal(&response)
	if err != nil {
		return fmt.Errorf("handleDeviceInfo: marshalling `DeviceInfoResponse`: %w", err)
	}

	err = frames.Write(data, protobuf.DeviceInfoResponseType, conn)
	if err != nil {
		return fmt.Errorf("failed sending DeviceInfoResponse: %w", err)
	}
	return nil
}
