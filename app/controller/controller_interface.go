package controller

import "github.com/gofiber/fiber/v2"

type Controller interface {
	RenderDict(c *fiber.Ctx) error
	RenderMain(c *fiber.Ctx) error
	RenderWord(c *fiber.Ctx) error
	RenderTags(c *fiber.Ctx) error
	RenderTerm(c *fiber.Ctx) error
	Search(c *fiber.Ctx) error
	SitemapGenerate(c *fiber.Ctx) error
}
