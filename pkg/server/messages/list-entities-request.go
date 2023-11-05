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

type ListEntities struct {
	err     error
	payload *api.ListEntitiesRequest
}

func ListEntitiesRequest(frame *frames.Frame) ListEntities {
	// parse frame into listEntities object
	h := ListEntities{
		payload: &api.ListEntitiesRequest{},
	}
	h.err = proto.Unmarshal(frame.Data, h.payload)
	return h
}

func (H ListEntities) Respond(conn net.Conn, config conf.Config) error {
	if H.err != nil {
		return fmt.Errorf("failed parsing ListEntitiesRequest frame: %w", H.err)
	}

	// FIXME: loop over sensors

	response := api.ListEntitiesDoneResponse{}
	data, err := proto.Marshal(&response)
	if err != nil {
		return fmt.Errorf("handleListEntities: marshalling `ListEntitiesResponse`: %w", err)
	}

	err = frames.Write(data, protobuf.ListEntitiesDoneResponseType, conn)
	if err != nil {
		return fmt.Errorf("failed sending ListEntitiesResponse: %w", err)
	}
	return nil
}
