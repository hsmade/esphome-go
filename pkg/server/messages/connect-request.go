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

type Connect struct {
	err     error
	payload *api.ConnectRequest
}

func ConnectRequest(frame *frames.Frame) Connect {
	// parse frame into connect object
	h := Connect{
		payload: &api.ConnectRequest{},
	}
	h.err = proto.Unmarshal(frame.Data, h.payload)
	return h
}

func (H Connect) Respond(conn net.Conn, config conf.Config) error {
	if H.err != nil {
		return fmt.Errorf("failed parsing ConnectRequest frame: %w", H.err)
	}

	response := api.ConnectResponse{
		InvalidPassword: !config.VerifyPasswordCallback(H.payload.Password),
	}
	data, err := proto.Marshal(&response)
	if err != nil {
		return fmt.Errorf("handleConnect: marshalling `ConnectResponse`: %w", err)
	}

	err = frames.Write(data, protobuf.ConnectResponseType, conn)
	if err != nil {
		return fmt.Errorf("failed sending ConnectResponse: %w", err)
	}
	return nil
}
