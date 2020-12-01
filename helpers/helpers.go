package helpers

import (
	"crypto/sha1"
	"encoding/binary"
	"encoding/hex"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"os"
)

func MarshalJSON(msg proto.Message) ([]byte, error) {
	return protojson.Marshal(msg)
}

func UnmarshalJSON(bits []byte, msg proto.Message) error {
	return protojson.Unmarshal(bits, msg)
}

func EnvOr(key string, defaul string) string {
	if val := os.Getenv(key); val == "" {
		return defaul
	} else {
		return val
	}
}

func Hash(val []byte) string {
	h := sha1.New()
	h.Write(val)
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}

func Uint64ToBytes(i uint64) []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	binary.PutUvarint(buf, i)
	return buf
}

func BytesToUint64(data []byte) uint64 {
	if len(data) == 0 {
		return 0
	}
	return binary.BigEndian.Uint64(data)
}
