package events

import (
	"log"
)

func (e *Events) OnGameStart() error {
	log.Println("Game start")
	return nil //if err isn't nil, HandleGame() will return it.
}
