package controllers

import (
	"fmt"
	"github.com/app-dictionary/app/helpers"
	"github.com/app-dictionary/app/services"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type Meta struct {
	Site            string
	SiteDescription string
	Host            string
	MetaTitle       string
	MetaDescription string
}

var meta = Meta{
	Site:            "Encycloped.ru",
	SiteDescription: "Библиотека словарей",
	Host:            "https://encycloped.ru/",
	MetaTitle:       "Словари, энциклопедии и справочники",
	MetaDescription: "Словари онлайн – один из самых популярных, наполненных и общедоступных ресурсов. Пользователь здесь найдет полезную информацию, затрагивающую все сферы человеческой деятельности, развития, культуры, языков и не только.",
}

func RenderMain(c *fiber.Ctx) error {
	dictionaries := services.AllDictionaries()

	//dictionaries = dictionaries[:10]

	columns := services.GetChunks(dictionaries, 2)

	return c.Render("views/main", fiber.Map{
		"Site":            meta.Site,
		"SiteDescription": meta.SiteDescription,
		"Host":            meta.Host,
		"MetaTitle":       meta.MetaTitle,
		"MetaDescription": meta.MetaDescription,
		"Title":           "Словари, энциклопедии и справочники",
		"Columns":         columns,
	})
}

func RenderTags(c *fiber.Ctx) error {
	return c.Render("tags", fiber.Map{
		"Site":            meta.Site,
		"SiteDescription": meta.SiteDescription,
		"Host":            meta.Host,
		"MetaTitle":       meta.MetaTitle,
		"MetaDescription": meta.MetaDescription,
		"title":           "TAGS",
	})
}

func RenderDictionary(c *fiber.Ctx) error {
	slug := c.Params("dictionary")
	letter := c.Query("letter", "")
	page := c.QueryInt("page", 1)

	limit := 50
	offset := 0
	if page == 2 {
		offset = limit
	} else if page > 2 {
		offset = (page - 1) * limit
	}

	dictionary := services.GetDictionary(slug)
	if dictionary.ID == 0 {
		return c.Redirect("https://encycloped.ru", 301)
	}

	letters := services.GetDictionaryLetters(dictionary.ID)
	if letter == "" && len(letters) > 0 {
		letter = letters[0].Title
	}

	letterModel := services.GetLetter(dictionary.ID, letter)
	words := services.GetWordsByDictionaryID(dictionary.ID, letterModel.ID, limit, offset)

	var prev int
	var next int

	if page > 0 {
		prev = page - 1
	} else {
		prev = 0
	}
	next = page + 1

	for i, l := range letters {
		letters[i].IsActive = l.Title == letter
	}

	return c.Render("views/dictionary", fiber.Map{
		"Site":            meta.Site,
		"SiteDescription": meta.SiteDescription,
		"Host":            meta.Host,
		"MetaTitle":       dictionary.Title,
		"MetaDescription": dictionary.Title,

		//"title":       c.Params("dictionary"),
		"dictionary":  dictionary,
		"letters":     letters,
		"letter":      letter,
		"words":       words,
		"page":        page,
		"prev":        prev,
		"next":        next,
		"canPrevPage": page > 1,
		"canNextPage": len(words) == limit,
	})
}

func RenderWord(c *fiber.Ctx) error {
	slugWord := c.Params("word")
	slugDict := c.Params("dictionary")

	dictionary := services.GetDictionaryBySlug(slugDict)
	word := services.GetWord(dictionary.ID, slugWord)
	if word.ID == 0 {
		return c.Redirect("https://encycloped.ru", 301)
	}

	words := services.GetWordsByTitle(dictionary.ID, word.Title)
	for i, w := range words {
		words[i].Title = strings.ToLower(w.Title)
		words[i].Content = fmt.Sprintf("%s...", helpers.ClearText(w.Content, 170))
	}

	prevLink, nextLink := services.GetWordNav(word.ID)

	return c.Render("views/word", fiber.Map{
		"Site":            meta.Site,
		"SiteDescription": meta.SiteDescription,
		"Host":            meta.Host,
		"MetaTitle":       word.Title + " - " + dictionary.Title,
		"MetaDescription": word.Title + " - " + dictionary.Title,

		"title":    word.Title,
		"word":     word,
		"words":    words,
		"prevLink": prevLink,
		"nextLink": nextLink,
	})
}
