package events

import (
	"log"
	"time"
)

func (e *Events) OnDeath() error {
	log.Println("Died and Respawned")
	// If we exclude Respawn(...) then the player won't press the "Respawn" button upon death
	go func() {
		time.Sleep(time.Second * 5)
		err := e.Player.Respawn()
		if err != nil {
			log.Print(err)
		}
	}()
	return nil
}
