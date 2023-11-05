//go:generate protoc --go_out=.. --go_opt=module=github.com/hsmade/esphome-go --go_opt=Mapi.proto=github.com/hsmade/esphome-go/proto/api --go_opt=Mapi_options.proto=github.com/hsmade/esphome-go/proto/api/options api.proto api_options.proto
package protobuf

import "encoding/binary"

type ProtoIntVar uint64

func ParseProtoVarInt(data []byte, size int) ProtoIntVar {
	var result uint64
	var bitpos uint8
	for i := 0; i < size; i++ {
		var val uint8 = data[i]
		result |= uint64(val&0x7F) << uint64(bitpos)
		bitpos += 7
		if (val & 0x80) == 0 {
			return ProtoIntVar(result)
		}
	}
	return ProtoIntVar(0)
}

func (P ProtoIntVar) Encode() []byte {
	val := uint64(P)
	result := make([]byte, 8)

	if val <= 0x7F {
		binary.LittleEndian.PutUint64(result, val)
		return result[:1]
	}

	size := 1
	for {
		if val <= 0 {
			break
		}
		temp := uint8(val & 0x7F)
		val >>= 7
		if val > 0 {
			result = append(result, temp|0x80)
		} else {
			result = append(result, temp)
		}
		size++
	}
	return result[:size]
}
