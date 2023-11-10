package helpers

import (
	"fmt"
	"github.com/sabloger/sitemap-generator/smg"
	"log"
	"regexp"
	"strings"
	"time"
)

func SiteMapGenerator(links []string) {
	// Генерация карты сайта
	sitemapLinks := GetChunksStr(links, GetDenominator(len(links), 50000))
	FileCreateSiteMap(true, sitemapLinks)
}

func GetDenominator(c, i int) int {
	r := c / i
	if r == 0 {
		return 1
	}
	return r
}

func GetChunksStr(page []string, size int) [][]string {
	var chunk [][]string
	chunkSize := (len(page) + size - 1) / size
	for i := 0; i < len(page); i += chunkSize {
		end := i + chunkSize
		if end > len(page) {
			end = len(page)
		}
		chunk = append(chunk, page[i:end])
	}

	return chunk
}
func FileCreateSiteMap(compress bool, urls [][]string) {
	now := time.Now().UTC()
	smi := smg.NewSitemapIndex(true)
	smi.SetCompress(compress)
	smi.SetHostname("https://encycloped.ru/")
	smi.SetOutputPath("public/sitemaps")

	// TODO: предусмотреть ограничения 50 000 ссылок и 50 mb на файл.
	for _, arr := range urls {
		smBlog := smi.NewSitemap()
		smBlog.SetLastMod(&now)
		for _, url := range arr {
			err := smBlog.Add(&smg.SitemapLoc{
				Loc:        url,
				LastMod:    &now,
				ChangeFreq: smg.Weekly,
				Priority:   0.8,
			})
			if err != nil {
				fmt.Println("err", err)
				log.Fatal(err)
			}
		}
	}

	_, err := smi.Save()
	if err != nil {
		fmt.Println("err", err)
		log.Fatal(err)
	}

	//fmt.Println("res", res)
}

func StringClearDataAttribute(s string) string {

	allHTML := regexp.MustCompile(`<([a-z]+) *[^/]*?>`)
	s = allHTML.ReplaceAllString(s, ``)

	emptyHTML := regexp.MustCompile(`<[\S]+><\/[\S]+>`)
	s = emptyHTML.ReplaceAllString(s, "")

	s = strings.ReplaceAll(s, "                      ", " ")
	s = strings.ReplaceAll(s, "                     ", " ")
	s = strings.ReplaceAll(s, "                    ", " ")
	s = strings.ReplaceAll(s, "                   ", " ")
	s = strings.ReplaceAll(s, "                  ", " ")
	s = strings.ReplaceAll(s, "                 ", " ")
	s = strings.ReplaceAll(s, "                ", " ")
	s = strings.ReplaceAll(s, "               ", " ")
	s = strings.ReplaceAll(s, "              ", " ")
	s = strings.ReplaceAll(s, "             ", " ")
	s = strings.ReplaceAll(s, "            ", " ")
	s = strings.ReplaceAll(s, "           ", " ")
	s = strings.ReplaceAll(s, "          ", " ")
	s = strings.ReplaceAll(s, "         ", " ")
	s = strings.ReplaceAll(s, "        ", " ")
	s = strings.ReplaceAll(s, "       ", " ")
	s = strings.ReplaceAll(s, "      ", " ")
	s = strings.ReplaceAll(s, "     ", " ")
	s = strings.ReplaceAll(s, "    ", " ")
	s = strings.ReplaceAll(s, "   ", " ")
	s = strings.ReplaceAll(s, "  ", " ")

	s = strings.ReplaceAll(s, "<p>", "")
	s = strings.ReplaceAll(s, "<div>", "")
	s = strings.ReplaceAll(s, "<span>", "")
	s = strings.ReplaceAll(s, "<small>", "")
	s = strings.ReplaceAll(s, "<a>", "")
	s = strings.ReplaceAll(s, "<ul>", "")
	s = strings.ReplaceAll(s, "<ol>", "")
	s = strings.ReplaceAll(s, "<li>", "")
	s = strings.ReplaceAll(s, "<img>", "")
	s = strings.ReplaceAll(s, "<img/>", "")
	s = strings.ReplaceAll(s, "<b>", "")
	s = strings.ReplaceAll(s, "<dt>", "")
	s = strings.ReplaceAll(s, "<tr>", "")
	s = strings.ReplaceAll(s, "<td>", "")
	s = strings.ReplaceAll(s, "<br>", "")
	s = strings.ReplaceAll(s, "</p>", "")
	s = strings.ReplaceAll(s, "</div>", "")
	s = strings.ReplaceAll(s, "</span>", "")
	s = strings.ReplaceAll(s, "</small>", "")
	s = strings.ReplaceAll(s, "</a>", "")
	s = strings.ReplaceAll(s, "</ul>", "")
	s = strings.ReplaceAll(s, "</ol>", "")
	s = strings.ReplaceAll(s, "</li>", "")
	s = strings.ReplaceAll(s, "</b>", "")
	s = strings.ReplaceAll(s, "</dt>", "")
	s = strings.ReplaceAll(s, "</tr>", "")
	s = strings.ReplaceAll(s, "</td>", "")
	s = strings.ReplaceAll(s, "\t", " ")
	s = strings.ReplaceAll(s, "\t", " ")
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, " ", " ")

	unicodeHTML := regexp.MustCompile(`[^A-zА-я !\.]`)
	s = unicodeHTML.ReplaceAllString(s, ``)

	s = strings.TrimSpace(s)

	regexpAttributeTags := regexp.MustCompile(`<([a-z]+\b)([^>]+?)\s?\/?\>`)
	s = regexpAttributeTags.ReplaceAllString(s, `<$1>`)

	regexpLinks := regexp.MustCompile(`<\/?a[^>]*>`)
	s = regexpLinks.ReplaceAllString(s, ``)

	regexpImages := regexp.MustCompile(`<\/?img[^>]*>`)
	s = regexpImages.ReplaceAllString(s, ``)

	s = strings.ReplaceAll(s, "......", ".")
	s = strings.ReplaceAll(s, ".....", ".")
	s = strings.ReplaceAll(s, "....", ".")
	s = strings.ReplaceAll(s, "...", ".")
	s = strings.ReplaceAll(s, "..", ".")

	return s
}

func cut(text string, limit int) string {
	runes := []rune(text)
	if len(runes) >= limit {
		return string(runes[:limit])
	}
	return text
}

func ClearText(s string, max int) string {
	ss := StringClearDataAttribute(s)

	if len(ss) < max {
		return ss
	}

	return cut(ss, max) //ss[:max]
}

func ToLover(s string) string {
	s = strings.ReplaceAll(s, " ", "-")
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, ":", "")
	s = strings.ReplaceAll(s, ";", "")
	s = strings.ReplaceAll(s, "(", "")
	s = strings.ReplaceAll(s, "{", "")
	s = strings.ReplaceAll(s, "[", "")
	s = strings.ReplaceAll(s, "|", "")
	s = strings.ReplaceAll(s, "\\", "")
	s = strings.ReplaceAll(s, "/", "")
	s = strings.ReplaceAll(s, ")", "")
	s = strings.ReplaceAll(s, "}", "")
	s = strings.ReplaceAll(s, "]", "")
	s = strings.ReplaceAll(s, ">", "")
	s = strings.ReplaceAll(s, "<", "")
	s = strings.ReplaceAll(s, "*", "")
	s = strings.ReplaceAll(s, "%", "")
	s = strings.ReplaceAll(s, "$", "")
	s = strings.ReplaceAll(s, "#", "")
	s = strings.ReplaceAll(s, "@", "")
	s = strings.ReplaceAll(s, "'", "")
	s = strings.ReplaceAll(s, "\"", "")
	s = strings.ReplaceAll(s, "`", "")
	s = strings.ReplaceAll(s, "+", "")
	s = strings.ReplaceAll(s, "&", "")
	s = strings.ReplaceAll(s, "!", "")
	s = strings.ReplaceAll(s, "?", "")
	s = strings.ReplaceAll(s, "~", "")
	s = strings.ReplaceAll(s, "№", "")
	s = strings.TrimSpace(s)
	return strings.ToLower(s)
}
