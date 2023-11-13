package controller

import (
	"errors"
	"github.com/app-dictionary/app/helpers"
	"github.com/gofiber/fiber/v2"
)

func (ctr *AppController) RenderDict(c *fiber.Ctx) error {
	slug := c.Params("dictionary")
	data, err := helpers.GetStaticFiles(ctr.EmbedFS, slug)
	if err == nil && data != nil {
		return c.Send(data)
	}

	letter := c.Query("letter", "")
	page := c.QueryInt("page", 1)

	limit := 50
	offset := 0
	if page == 2 {
		offset = limit
	} else if page > 2 {
		offset = (page - 1) * limit
	}

	dictionary, err := ctr.DictRepository.GetBySlug(slug)
	if err != nil {
		return c.Render("views/errors/500", ctr.GetError(err))
	}

	if dictionary.ID == 0 {
		return c.Render("views/errors/400", ctr.GetError(errors.New("словарь не найден")))
	}

	letters, err := ctr.LetterRepository.GetAllByDict(dictionary.ID)
	if err != nil {
		return c.Render("views/errors/500", ctr.GetError(err))
	}

	if letter == "" && len(letters) > 0 {
		letter = letters[0].Title
	}

	letterModel, err := ctr.LetterRepository.GetByDict(dictionary.ID, letter)
	if err != nil {
		return c.Render("views/errors/500", ctr.GetError(err))
	}

	words, err := ctr.WordRepository.GetByDictAndLetter(dictionary.ID, letterModel.ID, limit, offset)
	if err != nil {
		return c.Render("views/errors/500", ctr.GetError(err))
	}

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

	ctr.Data.Dictionary = dictionary
	ctr.Data.Letters = letters
	ctr.Data.Letter = letter
	ctr.Data.DictionaryWords = words
	ctr.Data.Page = page
	ctr.Data.Prev = prev
	ctr.Data.Next = next
	ctr.Data.CanPrevPage = page > 1
	ctr.Data.CanNextPage = len(words) == limit

	return c.Render("views/dictionary", ctr.GetResponse())
}
