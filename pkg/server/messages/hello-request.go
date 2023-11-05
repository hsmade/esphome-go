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

type Hello struct {
	err     error
	payload *api.HelloRequest
}

func HelloRequest(frame *frames.Frame) Hello {
	// parse frame into hello object
	h := Hello{
		payload: &api.HelloRequest{},
	}
	h.err = proto.Unmarshal(frame.Data, h.payload)
	return h
}

func (H Hello) Respond(conn net.Conn, config conf.Config) error {
	if H.err != nil {
		return fmt.Errorf("failed parsing HelloRequest frame: %w", H.err)
	}

	// FIXME: do something with the version?
	response := api.HelloResponse{
		ApiVersionMajor: H.payload.ApiVersionMajor,
		ApiVersionMinor: H.payload.ApiVersionMinor,
		ServerInfo:      "esphome-go",
		Name:            config.DeviceInfo.Name,
	}
	data, err := proto.Marshal(&response)
	if err != nil {
		return fmt.Errorf("handleHello: marshalling `HelloResponse`: %w", err)
	}

	err = frames.Write(data, protobuf.HelloResponseType, conn)
	if err != nil {
		return fmt.Errorf("failed sending HelloResponse: %w", err)
	}
	return nil
}
