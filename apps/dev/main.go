package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lavren1974/lavren1974cms/core/config"
)

func main() {

	cfg, err := config.LoadConfig("./config/default.toml")
	if err != nil {
		log.Fatal(err)
	}

	cfgGlobal, err := config.LoadConfigGlobal()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cfgGlobal)
	// LoadConfigGlobal

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(cfg.AppTitle)
	})

	app.Listen(cfg.AppPort)
}
