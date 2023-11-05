package frames

import "C"
import (
	"fmt"
	"github.com/hsmade/esphome-go/protobuf"
	"log/slog"
	"net"
	"time"
)

type Frame struct {
	MsgType protobuf.MsgType
	Size    uint32
	Data    []byte
}

func Read(conn net.Conn) (*Frame, error) {
	msg := Frame{}

	// Plain text implementation only
	// read header
	buffer := make([]byte, 1)
	n, err := conn.Read(buffer)
	if err != nil {
		return nil, fmt.Errorf("Read: reading header from connection: %w", err)
	}
	slog.Debug("Read: reading header", "Size", n, "Data", fmt.Sprintf("%02x", buffer))
	if buffer[0] != 0x00 {
		return nil, fmt.Errorf("Read: invalid header received: %02x, only plain text is implemented", buffer[0])
	}

	// read Size
	_ = conn.SetReadDeadline(time.Now().Add(time.Second))
	n, err = conn.Read(buffer)
	if err != nil {
		return nil, fmt.Errorf("Read: reading Size from connection: %w", err)
	}
	slog.Debug("Read: reading Size", "Size", n, "Data", fmt.Sprintf("%02x", buffer))
	msg.Size = uint32(protobuf.ParseProtoVarInt(buffer, 1))

	// read type
	_ = conn.SetReadDeadline(time.Now().Add(time.Second))
	n, err = conn.Read(buffer)
	if err != nil {
		return nil, fmt.Errorf("Read: reading type from connection: %w", err)
	}
	slog.Debug("Read: reading type", "Size", n, "Data", fmt.Sprintf("%02x", buffer))
	msg.MsgType = protobuf.MsgType(protobuf.ParseProtoVarInt(buffer, 1))

	// next, Data is read using Size

	buffer = make([]byte, msg.Size)
	n, err = conn.Read(buffer)
	if err != nil {
		return nil, fmt.Errorf("Read: reading Data from connection: %w", err)
	}
	slog.Debug("Read: received Data", "Size", n, "Data", buffer)
	msg.Data = buffer
	return &msg, nil
}

func Write(msg []byte, msgType protobuf.MsgType, conn net.Conn) error {
	slog.Debug("Write", "msg", fmt.Sprintf("%02x", msg), "type", msgType)
	slog.Debug("Write: sending header", "Data", fmt.Sprintf("%02x", 0x00))
	_, err := conn.Write([]byte{0x00})
	if err != nil {
		return fmt.Errorf("Write: writing header: %w", err)
	}

	data := protobuf.ProtoIntVar(len(msg)).Encode()
	slog.Debug("Write: sending size", "Data", fmt.Sprintf("%02x", data))
	_, err = conn.Write(data)
	if err != nil {
		return fmt.Errorf("Write: writing Size: %w", err)
	}

	data = protobuf.ProtoIntVar(msgType).Encode()
	slog.Debug("Write: sending type", "Data", fmt.Sprintf("%02x", data))
	_, err = conn.Write(data)
	if err != nil {
		return fmt.Errorf("Write: writing type: %w", err)
	}

	_, err = conn.Write(msg)
	if err != nil {
		return fmt.Errorf("Write: writing Data: %w", err)
	}

	return nil
}
