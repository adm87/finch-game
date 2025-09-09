package main

import (
	"github.com/adm87/finch-application/application"
	"github.com/adm87/finch-game/game"
)

func main() {
	cmd := application.NewApplicationCommand("finch-game", game.Game)
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
