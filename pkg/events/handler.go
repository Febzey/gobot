package events

import (
	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/bot/basic"
)

type Events struct {
	Player *basic.Player
}

func (e *Events) RegisterEvents(client *bot.Client) {

	player := basic.NewPlayer(client, basic.DefaultSettings)
	e.Player = player

	basic.EventsListener{
		GameStart: e.OnGameStart,
		// ChatMsg:    e.OnChatMsg,
		// SystemMsg:  e.OnSystemMsg,
		Disconnect: e.OnDisconnect,
		Death:      e.OnDeath,
	}.Attach(client)
}