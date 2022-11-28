package events

import (
	"log"

	pk "github.com/Tnze/go-mc/net/packet"
)

func (e *Events) OnLogin(_ pk.Packet) error {
	log.Println("Game start")
	return nil
}
