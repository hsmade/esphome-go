package frames

import "C"
import (
	"fmt"
	"github.com/hsmade/esphome-go/protobuf"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"log/slog"
	"net"
	"time"
)

type Frame struct {
	MsgType protobuf.MsgType
	Size    uint32
	Data    []byte
}

var (
	promFramesReceived = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "esphome_go_frames_received_total",
		Help: "The total amount of frames received",
	},
		[]string{"message_type"},
	)

	promFrameReadFailuresTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "esphome_go_frame_read_failures_total",
		Help: "The total number of failures when reading frames",
	})

	promFramesSent = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "esphome_go_frames_sent_total",
		Help: "The total amount of frames sent",
	},
		[]string{"message_type"},
	)

	promFrameWriteFailuresTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "esphome_go_frame_write_failures_total",
		Help: "The total number of failures when writing frames",
	})
)

func Read(conn net.Conn) (*Frame, error) {
	msg := Frame{}

	// Plain text implementation only
	// read header
	buffer := make([]byte, 1)
	n, err := conn.Read(buffer)
	if err != nil {
		return nil, fmt.Errorf("Read: reading header from connection: %w", err)
	}
	slog.Debug("Read: reading header", "Size", n, "data", fmt.Sprintf("%02x", buffer))
	if buffer[0] != 0x00 {
		promFrameReadFailuresTotal.Inc()
		return nil, fmt.Errorf("Read: invalid header received: %02x, only plain text is implemented", buffer[0])
	}

	// read Size
	_ = conn.SetReadDeadline(time.Now().Add(time.Second))
	n, err = conn.Read(buffer)
	if err != nil {
		promFrameReadFailuresTotal.Inc()
		return nil, fmt.Errorf("Read: reading Size from connection: %w", err)
	}
	slog.Debug("Read: reading Size", "Size", n, "data", fmt.Sprintf("%02x", buffer))
	msg.Size = uint32(protobuf.ParseProtoVarInt(buffer, 1))

	// read type
	_ = conn.SetReadDeadline(time.Now().Add(time.Second))
	n, err = conn.Read(buffer)
	if err != nil {
		promFrameReadFailuresTotal.Inc()
		return nil, fmt.Errorf("Read: reading type from connection: %w", err)
	}
	msg.MsgType = protobuf.MsgType(protobuf.ParseProtoVarInt(buffer, 1))
	slog.Debug("Read: reading type", "Size", n, "data", fmt.Sprintf("%02x", buffer), "type", msg.MsgType.String())
	promFramesReceived.WithLabelValues(msg.MsgType.String()).Inc()

	// next, Data is read using Size

	buffer = make([]byte, msg.Size)
	n, err = conn.Read(buffer)
	if err != nil {
		promFrameReadFailuresTotal.Inc()
		return nil, fmt.Errorf("Read: reading Data from connection: %w", err)
	}
	slog.Debug("Read: received Data", "Size", n, "data", buffer)
	msg.Data = buffer
	return &msg, nil
}

func Write(msg []byte, msgType protobuf.MsgType, conn net.Conn) error {
	slog.Debug("Write", "msg", fmt.Sprintf("%02x", msg), "type", msgType.String())
	slog.Debug("Write: sending header", "data", fmt.Sprintf("%02x", 0x00))
	_, err := conn.Write([]byte{0x00})
	if err != nil {
		promFrameWriteFailuresTotal.Inc()
		return fmt.Errorf("Write: writing header: %w", err)
	}

	data := protobuf.ProtoIntVar(len(msg)).Encode()
	slog.Debug("Write: sending size", "data", fmt.Sprintf("%02x", data))
	_, err = conn.Write(data)
	if err != nil {
		promFrameWriteFailuresTotal.Inc()
		return fmt.Errorf("Write: writing Size: %w", err)
	}

	data = protobuf.ProtoIntVar(msgType).Encode()
	slog.Debug("Write: sending type", "data", fmt.Sprintf("%02x", data), "type", msgType)
	_, err = conn.Write(data)
	if err != nil {
		promFrameWriteFailuresTotal.Inc()
		return fmt.Errorf("Write: writing type: %w", err)
	}
	promFramesSent.WithLabelValues(msgType.String()).Inc()

	_, err = conn.Write(msg)
	if err != nil {
		promFrameWriteFailuresTotal.Inc()
		return fmt.Errorf("Write: writing Data: %w", err)
	}

	return nil
}
