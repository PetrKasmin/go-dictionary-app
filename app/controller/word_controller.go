package controller

import (
	"fmt"
	"github.com/app-dictionary/app/helpers"
	"github.com/app-dictionary/app/repositories"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func (ctr *AppController) RenderWord(c *fiber.Ctx) error {
	slugWord := c.Params("word")
	data, err := helpers.GetStaticFiles(ctr.EmbedFS, slugWord)
	if err == nil && data != nil {
		return c.Send(data)
	}

	slugDict := c.Params("dictionary")

	dictionary, err := ctr.DictRepository.GetBySlug(slugDict)
	if err != nil {
		return c.Render("views/errors/500", ctr.GetError(err))
	}

	word, err := ctr.WordRepository.GetWordByDictAndTitle(dictionary.ID, slugWord)
	if err != nil {
		return c.Render("views/errors/500", ctr.GetError(err))
	}

	if word.ID == 0 {
		return c.Render("views/errors/400", ctr.GetError(err))
	}

	words, err := ctr.WordRepository.GetWordsByDictAndTitle(dictionary.ID, word.Title)
	if err != nil {
		return c.Render("views/errors/500", ctr.GetError(err))
	}

	for i, w := range words {
		words[i].Title = strings.ToLower(w.Title)
		words[i].Content = fmt.Sprintf("%s...", helpers.ClearText(w.Content, 170))
	}

	wordsForNav, err := ctr.WordRepository.GetWordForNav(word.ID)
	if err != nil {
		return c.Render("views/errors/500", ctr.GetError(err))
	}

	var prev, next repositories.Link

	for _, w := range wordsForNav {
		if w.ID < word.ID {
			prev.Title = w.Title
			prev.Slug = w.Slug
			prev.DictionarySlug = w.DictionarySlug
		} else if w.ID > word.ID {
			next.Title = w.Title
			next.Slug = w.Slug
			next.DictionarySlug = w.DictionarySlug
		}
	}

	ctr.Data.Title = word.Title
	ctr.Data.Words = words
	ctr.Data.Word = word
	ctr.Data.PrevLink = prev
	ctr.Data.NextLink = next

	return c.Render("views/word", ctr.GetResponse())
}