package script

import "github.com/gofiber/fiber/v2"

type ScriptHandler struct {
	name string
	path string
}

func New(path string) ScriptHandler {
	return ScriptHandler{
		name: "Script event handler",
		path: path,
	}
}

func (sh ScriptHandler) GetName() string {
	return sh.name
}

func (sh ScriptHandler) Handle(c *fiber.Ctx) error {
	return nil
}
