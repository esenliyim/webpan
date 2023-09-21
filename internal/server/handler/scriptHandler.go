package handler

import (
	"log"

	"github.com/esenliyim/sp-tray/internal/server/handler/script"
)

func init() {
	log.Println("registering script event handler")
	RegisterHandler("Webpan-Script", func(args map[string]string) Handler {
		return script.New(args["path"])
	})
}
