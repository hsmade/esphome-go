package server

//
//import (
//	"github.com/hsmade/esphome-go/protobuf/api"
//	"google.golang.org/protobuf/proto"
//	"log/slog"
//	"time"
//)
//
//func (C *Connection) Pinger() {
//	ticker := time.NewTicker(10 * time.Second)
//	for {
//		select {
//		case <-ticker.C:
//			if time.Now().Sub(C.lastReceived).Seconds() < 60 {
//				continue
//			}
//			msg := api.PingRequest{}
//			data, err := proto.Marshal(&msg)
//			if err != nil {
//				slog.Warn("Pinger: marshalling ping: %w", err)
//			}
//			err = C.SendFrame(data, 7)
//			if err != nil {
//				slog.Warn("Pinger: sending ping: %w", err)
//			}
//			return
//		}
//	}
//}
