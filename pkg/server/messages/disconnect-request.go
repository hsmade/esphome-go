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

type Disconnect struct {
	err     error
	payload *api.DisconnectRequest
}

func DisconnectRequest(frame *frames.Frame) Disconnect {
	// parse frame into disconnect object
	h := Disconnect{
		payload: &api.DisconnectRequest{},
	}
	h.err = proto.Unmarshal(frame.Data, h.payload)
	return h
}

func (H Disconnect) Respond(conn net.Conn, config conf.Config) error {
	if H.err != nil {
		return fmt.Errorf("failed parsing DisconnectRequest frame: %w", H.err)
	}

	response := api.DisconnectResponse{}
	data, err := proto.Marshal(&response)
	if err != nil {
		return fmt.Errorf("handleDisconnect: marshalling `DisconnectResponse`: %w", err)
	}

	err = frames.Write(data, protobuf.DisconnectResponseType, conn)
	if err != nil {
		return fmt.Errorf("failed sending DisconnectResponse: %w", err)
	}
	return nil
}
