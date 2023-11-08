package server

import (
	"errors"
	"fmt"
	"github.com/hsmade/esphome-go/pkg/server/conf"
	"github.com/hsmade/esphome-go/pkg/server/frames"
	"github.com/hsmade/esphome-go/pkg/server/messages"
	"github.com/hsmade/esphome-go/protobuf"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"io"
	"log/slog"
	"net"
	"os"
	"time"
)

var (
	promConnectionsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "esphome_go_connections_total",
		Help: "The total number of connections handled",
	})

	promMessageHandleFailures = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "esphome_go_message_handle_failures_total",
		Help: "The total amount of failures when handling messages",
	},
		[]string{"message_type"},
	)
)

type Server struct {
	Port   int
	Config conf.Config
}

func (S *Server) Listen() error {
	slog.Info("starting server", "port", S.Port)
	l, err := net.Listen("tcp4", fmt.Sprintf(":%d", S.Port))
	if err != nil {
		return fmt.Errorf("starting listener: %w", err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			return fmt.Errorf("accepting connection: %w", err)
		}
		go S.handleConnection(conn)
	}
}

func (S *Server) handleConnection(conn net.Conn) {
	promConnectionsTotal.Inc()
	// FIXME: set up pinger, needs to know when last message was received
	for {
		_ = conn.SetReadDeadline(time.Now().Add(time.Second * 1))
		frame, err := frames.Read(conn)
		if err != nil {
			if errors.Is(err, io.EOF) {
				_ = conn.Close()
				return
			}
			if errors.Is(err, os.ErrDeadlineExceeded) {
				continue
			}
			slog.Error("failed reading frame: %w", err)
			continue
		}

		switch frame.MsgType {
		case protobuf.HelloRequestType:
			err = messages.HelloRequest(frame).Respond(conn, S.Config)
		case protobuf.ConnectRequestType:
			err = messages.ConnectRequest(frame).Respond(conn, S.Config)
		case protobuf.DisconnectRequestType:
			err = messages.DisconnectRequest(frame).Respond(conn, S.Config)
			_ = conn.Close()
			return
		case protobuf.PingRequestType:
			err = messages.PingRequest(frame).Respond(conn, S.Config)
		case protobuf.DeviceInfoRequestType:
			err = messages.DeviceInfoRequest(frame).Respond(conn, S.Config)
		case protobuf.ListEntitiesRequestType:
			err = messages.ListEntitiesRequest(frame).Respond(conn, S.Config)
		}
		if err != nil {
			slog.Error("failed handling message: %w", err)
			promMessageHandleFailures.WithLabelValues(frame.MsgType.String()).Inc()
			continue
		}

	}

}
