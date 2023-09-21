package main

import (
	"fmt"
	"os"

	"github.com/esenliyim/sp-tray/internal/server/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("port")
	if port == "" {
		port = "8080"
	}

	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})

	app.Post(fmt.Sprintf("/%s", os.Getenv("endpoint")), controller.ListenToHooks).Name("alsdkaskld")

	app.Listen(fmt.Sprintf("%s:%s", os.Getenv("host"), port))
}
