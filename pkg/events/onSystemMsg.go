package events

import (
	"fmt"
	"regexp"

	"github.com/Tnze/go-mc/chat"
)

// create a an array with different types of chat regex's
// var chatRegex = []string{
// 	"(?:[^ ]* )?([^ ]+)(?: [^ ]*)? » (.*)",
// 	"^<([^ ]*)> (.*)$",
// 	"^([^ ]*): (.*)$",
// }

func (e *Events) OnSystemMsg(c chat.Message, pos byte) error {
	msg := c.ClearString()

	fmt.Println("system:", msg)

	var User struct {
		Username string
		Message  string
		Uuid     string
		Ping     string
	}

	r, _ := regexp.Compile("<([^ ]*)> (.*)")
	if r.MatchString(msg) {
		matches := r.FindStringSubmatch(msg)
		User.Username = matches[1]
		User.Message = matches[2]

		fmt.Println("User: ", User.Username, "msg: ", User.Message)
	}

	// for _, regex := range chatRegex {
	// 	r, _ := regexp.Compile("^<([^ ]*)> (.*)$")
	// 	if r.MatchString(msg) {
	// 		matches := r.FindStringSubmatch(msg)
	// 		User.Username = matches[1]
	// 		User.Message = matches[2]
	// 	}
	// }

	return nil
}
