package github

import (
	"os"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func handlePush(c *fiber.Ctx) error {
	_, err := parsePayload(c)
	if err != nil {
		return err
	}
	_, err = exec.Command(os.Getenv("GITHUB_SCRIPT")).Output()
	if err != nil {
		c.SendStatus(500)
		return c.SendString(err.Error())
	}
	return c.SendString("push handled")
}

func parsePayload(c *fiber.Ctx) (*HookRequestPush, error) {
	p := new(HookRequestPush)

	if err := c.BodyParser(p); err != nil {
		return nil, err
	}

	return p, nil
}
