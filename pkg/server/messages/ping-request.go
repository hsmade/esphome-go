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

type Ping struct {
	err     error
	payload *api.PingRequest
}

func PingRequest(frame *frames.Frame) Ping {
	// parse frame into ping object
	h := Ping{
		payload: &api.PingRequest{},
	}
	h.err = proto.Unmarshal(frame.Data, h.payload)
	return h
}

func (H Ping) Respond(conn net.Conn, config conf.Config) error {
	if H.err != nil {
		return fmt.Errorf("failed parsing PingRequest frame: %w", H.err)
	}

	response := api.PingResponse{}
	data, err := proto.Marshal(&response)
	if err != nil {
		return fmt.Errorf("handlePing: marshalling `PingResponse`: %w", err)
	}

	err = frames.Write(data, protobuf.PingResponseType, conn)
	if err != nil {
		return fmt.Errorf("failed sending PingResponse: %w", err)
	}
	return nil
}
