package main

import (
	"log"

	"github.com/Tnze/go-mc/bot"
	GMMAuth "github.com/maxsupermanhd/go-mc-ms-auth"
)

func Authenticate() (bot.Auth, error) {
	mauth, er := GMMAuth.GetMCcredentials("./auth.cache", "88650e7e-efee-4857-b9a9-cf580a00ef43")
	if er != nil {
		return mauth, er
	}
	log.Print("Authenticated as ", mauth.Name, " (", mauth.UUID, ")")
	return mauth, nil
}
