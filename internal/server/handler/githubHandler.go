package handler

import (
	"log"

	"github.com/esenliyim/sp-tray/internal/server/handler/github"
)

func init() {
	log.Println("registering github event handler")
	RegisterHandler("GitHub-Hookshot", func(args map[string]string) Handler {
		return github.New(args["name"])
	})
}
