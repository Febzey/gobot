package events

import (
	"log"

	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/chat"
)

//create object to hold message and username
type User struct {
	Username string
	Message  string
}

func (e *Events) OnChatMsg(c *basic.PlayerMessage) error {

	sender := c.SenderDisplayName.ClearString()

	user := User{
		Username: sender,
		Message:  c.SignedMessage.ClearString(),
	}

	return handleMessages(user)

	// log.Printf("Chat: %v", c)s

	// // username := c.SenderDisplayName.String()
	// // message := c.SignedMessage.String()
	// // log.Println("Username:", username, "Chat:", message)

	// return nil
}

// create a an array with different types of chat regex's
// var chatRegex = []string{
// 	"(?:[^ ]* )?([^ ]+)(?: [^ ]*)? » (.*)",
// 	"^<([^ ]*)> (.*)$",
// 	"^([^ ]*): (.*)$",
// }

func filterMessages(msg string) error {

	return nil
}

func (e *Events) OnSystemMsg(c chat.Message, pos byte) error {

	user := User{
		Username: "systemm",
		Message:  c.ClearString(),
	}

	return handleMessages(user)
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
}

func handleMessages(user User) error {
	//destruct username and message from user
	username := user.Username
	message := user.Message

	log.Println("username:", username, "message:", message)
	return nil
}
