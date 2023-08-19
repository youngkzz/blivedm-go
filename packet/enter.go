package packet

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

type Enter struct {
	UID      int    `json:"uid"`
	RoomID   int    `json:"roomid"`
	ProtoVer int    `json:"protover"`
	Buvid    string `json:"buvid"`
	Platform string `json:"platform"`
	Type     int    `json:"type"`
	Key      string `json:"key"`
}

// NewEnterPacket 构造进入房间的包
// uid 可以为 0, key 在使用 broadcastlv 服务器的时候不需要
func NewEnterPacket(uid int, roomID int, key string) []byte {
	ent := &Enter{
		UID:      uid,
		RoomID:   roomID,
		ProtoVer: 3,
		Platform: "web",
		Type:     2,
		Key:      key,
	}
	m, err := json.Marshal(ent)
	if err != nil {
		log.Error("NewEnterPacket JsonMarshal failed", err)
	}
	pkt := NewPlainPacket(RoomEnter, m)
	return pkt.Build()
}
