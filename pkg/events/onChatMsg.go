package events

import (
	"log"

	"github.com/Tnze/go-mc/bot/basic"
)

func (e *Events) OnChatMsg(c *basic.PlayerMessage) error {

	log.Printf("Chat: %v", c)

	// username := c.SenderDisplayName.String()
	// message := c.SignedMessage.String()
	// log.Println("Username:", username, "Chat:", message)

	return nil
}
