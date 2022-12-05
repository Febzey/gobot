package events

import (
	"fmt"

	"github.com/Tnze/go-mc/chat"
	pk "github.com/Tnze/go-mc/net/packet"
)

// create object to hold message and username
type User struct {
	Username string
	Message  string
}

// func handleMessages(user User) error {
// 	username := user.Username
// 	message := user.Message

// 	log.Println("username:", username, "message:", message)
// 	return nil
// }

// func filterMessages(msg string) error {
// 	return nil
// }

// ClientboundPlayerChat
func (e *Events) OnPlayerMsg(p pk.Packet) error {

	var (
		SignedMessage chat.Message
		Unsigned      pk.Boolean
	)

	if err := p.Scan(&SignedMessage, &Unsigned); err != nil {
		return err
	}

	fmt.Println(SignedMessage)

	return nil
}

func (e *Events) OnSystemMsg(p pk.Packet) error {

	var (
		msg chat.Message
		pos pk.Boolean
	)

	if err := p.Scan(&msg, &pos); err != nil {
		return err
	}

	fmt.Println(msg)

	return nil
}

// user := User{
// 	Username: "systemm",
// 	Message:  c.ClearString(),
// }
// return handleMessages(user)
// msg := c.ClearString()

// fmt.Println("system:", msg)

// var User struct {
// 	Username string
// 	Message  string
// 	Uuid     string
// 	Ping     string
// }

// r, _ := regexp.Compile("<([^ ]*)> (.*)")
// if r.MatchString(msg) {
// 	matches := r.FindStringSubmatch(msg)
// 	User.Username = matches[1]
// 	User.Message = matches[2]

// 	fmt.Println("User: ", User.Username, "msg: ", User.Message)
// }

// // for _, regex := range chatRegex {
// // 	r, _ := regexp.Compile("^<([^ ]*)> (.*)$")
// // 	if r.MatchString(msg) {
// // 		matches := r.FindStringSubmatch(msg)
// // 		User.Username = matches[1]
// // 		User.Message = matches[2]
// // 	}
// // }
// func (e *Events) OnChatMsg(c *basic.PlayerMessage) error {

// 	sender := c.SenderDisplayName.ClearString()

// 	user := User{
// 		Username: sender,
// 		Message:  c.SignedMessage.ClearString(),
// 	}

// 	return handleMessages(user)

// 	// log.Printf("Chat: %v", c)s

// 	// // username := c.SenderDisplayName.String()
// 	// // message := c.SignedMessage.String()
// 	// // log.Println("Username:", username, "Chat:", message)

// 	// return nil
// }

// create a an array with different types of chat regex's
// var chatRegex = []string{
// 	"(?:[^ ]* )?([^ ]+)(?: [^ ]*)? » (.*)",
// 	"^<([^ ]*)> (.*)$",
// 	"^([^ ]*): (.*)$",
// }
