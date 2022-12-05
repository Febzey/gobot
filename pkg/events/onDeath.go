package events

import (
	"log"

	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

func (e *Events) HandleHealth(p pk.Packet) error {

	var health pk.Float
	var food pk.VarInt
	var foodSaturation pk.Float

	if err := p.Scan(&health, &food, &foodSaturation); err != nil {
		return err
	}

	if health == 0 {
		log.Println("Bot is dead")

		const PerformRespawn = 0

		err := e.Client.Conn.WritePacket(pk.Marshal(
			int32(packetid.ServerboundClientCommand),
			pk.VarInt(PerformRespawn),
		))
		if err != nil {
			return err
		}
	}

	return nil
}
