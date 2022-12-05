package events

import (
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

func (e *Events) HandleKeepAlive(p pk.Packet) error {
	var id pk.Long
	if err := p.Scan(&id); err != nil {
		return err
	}

	err := e.Client.Conn.WritePacket(pk.Packet{
		ID:   int32(packetid.ServerboundKeepAlive),
		Data: p.Data,
	})
	if err != nil {
		return err
	}

	return nil
}
