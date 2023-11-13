package controller

import "github.com/gofiber/fiber/v2"

type Response struct {
	Meta Meta
	Data Data
}

func (ctr *AppController) GetResponse() fiber.Map {
	return fiber.Map{
		"Meta": ctr.Meta,
		"Data": ctr.Data,
	}
}

func (ctr *AppController) GetError(err error) fiber.Map {
	return fiber.Map{
		"Error": err.Error(),
	}
}
