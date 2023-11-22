package helpers

import (
	"errors"
	"fmt"
	"github.com/app-dictionary/app/models"
	"github.com/sabloger/sitemap-generator/smg"
	"io"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode"
)

const (
	robotsFile        = "robots.txt"
	sitemapsDir       = "sitemaps"
	sitemapsIndexFile = "sitemap.xml.gz"
)

func SplitByLetter(dictionaries []models.Dictionary) ([]models.Dictionary, []string) {
	sort.Slice(dictionaries, func(i, j int) bool {
		return dictionaries[i].Title < dictionaries[j].Title
	})

	var letters []string
	var result []models.Dictionary
	firstDictionary := dictionaries[0]
	firstLetter := []rune(firstDictionary.Title)[0]

	for _, d := range dictionaries {
		currentLetter := []rune(d.Title)[0]
		if unicode.IsLetter(currentLetter) && firstLetter != currentLetter {
			firstLetter = currentLetter
			letters = append(letters, string(currentLetter))
			result = append(result, models.Dictionary{
				Title:     string(currentLetter),
				IsDivider: true,
			})
		}
		result = append(result, d)
	}

	return result, letters
}

func GetChunks(page []models.Dictionary, size int) [][]models.Dictionary {
	var chunk [][]models.Dictionary
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

func GetStaticFiles(embedFS http.FileSystem, param string) ([]byte, error) {
	var path string
	switch param {
	case sitemapsDir:
		path = fmt.Sprintf("public/%s/%s", param, sitemapsIndexFile)
	case sitemapsIndexFile:
		path = fmt.Sprintf("public/%s/%s", sitemapsDir, param)
	case robotsFile:
		path = fmt.Sprintf("public/%s", param)
	default:
		return nil, errors.New("no static")
	}

	file, err := embedFS.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func SitemapGenerator(links []string) {
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
