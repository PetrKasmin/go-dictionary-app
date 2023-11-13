package controller

type Meta struct {
	Site        string
	Host        string
	Title       string
	Description string
}

func NewMeta() Meta {
	title := "Словари, энциклопедии и справочники"

	description := "Словари онлайн – один из самых популярных, " +
		"наполненных и общедоступных ресурсов. Пользователь здесь найдет полезную информацию, " +
		"затрагивающую все сферы человеческой деятельности, развития, культуры, языков и не только."

	return Meta{
		Site:        "Encycloped.ru",
		Host:        "https://encycloped.ru/",
		Title:       title,
		Description: description,
	}
}
