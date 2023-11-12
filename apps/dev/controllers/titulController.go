package controllers

import (
	// "fiber-mvc/config"
	// "fiber-mvc/models"
	// "log"
	// "strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lavren1974/lavren1974cms/core/config"
)

type TitulController struct {
	Config *config.AppConfig
}

func NewTitulController(cfg *config.AppConfig) *TitulController {
	return &TitulController{Config: cfg}
}

func (controller *TitulController) GetTitul(c *fiber.Ctx) error {

	return c.Render("index", fiber.Map{
		"Title": controller.Config.AppTitle,
	})
}
