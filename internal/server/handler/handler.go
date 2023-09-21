package handler

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type HandlerFactory func(args map[string]string) Handler

type Handler interface {
	Handle(c *fiber.Ctx) error
	GetName() string
}

var (
	handlerRegistry = make(map[string]HandlerFactory)
)

func RegisterHandler(prefix string, factory HandlerFactory) error {
	if _, ok := handlerRegistry[prefix]; ok {
		return errors.New("a handler with the same prefix already exists")
	}
	handlerRegistry[prefix] = factory
	return nil
}

func Get(userAgent []byte) (HandlerFactory, error) {
	agentString := string(userAgent)

	for prefix, factory := range handlerRegistry {
		if strings.HasPrefix(agentString, prefix) {
			return factory, nil
		}
	}
	return nil, errors.New("could not match to a known handler")
}
