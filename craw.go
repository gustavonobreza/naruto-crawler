package main

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	ARTICLES_TO_TAKE = 7
)

// Fetch return (body, nil) if sucessful, and (nil, error) if not.
func Fetch(url, name string, txtSlince chan []string) {
	var atricles []string
	resp, err := http.Get(url)
	if err != nil {
		println("[ERROR]: Fail to fetch data from:", url)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		println("[ERROR]: Fail to fetch data from:", url)
		return
	}
	// Articles
	rawArticles := doc.Find("#mw-content-text > div > p").FilterFunction(func(i int, s *goquery.Selection) bool {
		t := s.Text()
		// Removing junk data
		return len(t) >= 10
	})

	rawTxtArticles := rawArticles.Text()

	// The page not respond with 404, so if not has "." the url fail
	if !strings.Contains(rawTxtArticles, ".") {
		println("[ERROR]: Fail to fetch data from:", url)
		return
	}

	// Removing junks
	rawTxtArticles = strings.ReplaceAll(rawTxtArticles, "\t", "")
	rawTxtArticles = strings.ReplaceAll(rawTxtArticles, "\r", "")

	// by '\n'
	splitted := strings.Split(rawTxtArticles, "\n")
	// by '.'
	// splitted := strings.Split(rawTxtArticles, ".")

	// Check if page not exists
	if strings.Contains(splitted[0], "There is currently no text in this page") {
		println("[ERROR]: Fail to fetch data from:", url)
		return
	}

	for i, v := range splitted {
		isSmall := len(strings.Join(atricles, "")) < 2500
		if i < ARTICLES_TO_TAKE || isSmall {
			txt := v
			txt = strings.TrimSpace(txt)

			// Removing junk as [23] or [7], I don't know what is it
			re := regexp.MustCompile(`(\[)+\d+(\])`)
			txtClean1 := re.ReplaceAllString(txt, "") + "."
			atricles = append(atricles, txtClean1)
		}
	}

	// println("SIZE OF", name, "->", len(atricles))
	// println("FIRST OF", name, "HAS", "->", len(atricles[0]))
	// println("TOTAL OF", name, "HAS", "->", len(strings.Join(atricles, "")))
	// println("--------------------------------------------")
	// print("\n")
	txtSlince <- atricles
}
