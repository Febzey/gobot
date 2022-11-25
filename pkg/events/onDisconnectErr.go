package events

import "github.com/Tnze/go-mc/chat"

type DisconnectErr struct {
	Reason chat.Message
}

func (d DisconnectErr) Error() string {
	return "disconnect: " + d.Reason.String()
}

func (e *Events) OnDisconnect(reason chat.Message) error {
	//return the discconect reason as a string
	return DisconnectErr{Reason: reason}
}
