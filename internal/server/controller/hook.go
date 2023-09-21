package controller

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"

	"github.com/esenliyim/sp-tray/internal/server/handler"
	"github.com/gofiber/fiber/v2"
)

func ListenToHooks(c *fiber.Ctx) error {

	// handler, err := handler.Get(c.Context().UserAgent())
	// if err != nil {
	// 	c.SendStatus(400)
	// 	return c.SendString(err.Error())
	// }
	loadListeners()
	return nil
	// return handler.Handle(c)
}

type Endpoint struct {
	Name    string          `yaml:"name"`
	Handler handler.Handler `yaml:"handler"`
	Post    Method          `yaml:"post"`
}

type Method struct {
	Name string `yaml:"name"`
	Cmd  string `yaml:"cmd"`
}

type Controller struct {
	Name      string     `yaml:"name"`
	Endpoints []Endpoint `yaml:"endpoints"`
}

func loadListeners() error {
	dir := "/Users/emre/sources/webpan/config"
	return filepath.WalkDir(dir, loadListener)
}

func loadListener(path string, file fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if !file.IsDir() {
		fmt.Println(path)
	}
	log.Println(path)

	return nil
}
