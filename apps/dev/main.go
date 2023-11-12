package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/lavren1974/lavren1974cms/apps/dev/controllers"
	"github.com/lavren1974/lavren1974cms/core/config"
)

func main() {

	cfg, err := config.LoadConfig("./config/default.toml")
	if err != nil {
		log.Fatal(err)
	}

	titulController := controllers.NewTitulController(cfg)

	cfgGlobal, err := config.LoadConfigGlobal()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cfgGlobal)
	// LoadConfigGlobal

	// Initialize the template engine and load all templates
	engine := html.New("./views", ".html")

	// Use ParseGlob to load all templates
	if err := engine.Load(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Get("/", titulController.GetTitul)

	app.Listen(cfg.AppPort)
}
