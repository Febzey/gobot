package events

import (
	"github.com/Tnze/go-mc/chat"
	pk "github.com/Tnze/go-mc/net/packet"
)

type DisconnectErr struct {
	Reason chat.Message
}

func (d DisconnectErr) Error() string {
	return "disconnect: " + d.Reason.String()
}

func (e *Events) HandleDisconnect(p pk.Packet) error {
	var reason chat.Message
	if err := p.Scan(&reason); err != nil {
		return DisconnectErr{Reason: reason}
	}

	return nil
}
