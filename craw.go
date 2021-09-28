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

var (
	atricles []string
)

// Fetch return (body, nil) if sucessful, and (nil, error) if not.
func Fetch(url string) ([]string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return []string{}, err
	}

	defer resp.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(resp.Body)

	rawArticles := doc.Find("#mw-content-text > div > p").FilterFunction(func(i int, s *goquery.Selection) bool {
		t := s.Text()
		// Removing junk data
		return len(t) >= 5
	})

	i := 0

	rawArticles.Each(func(_ int, s *goquery.Selection) {
		if i < 5 {
			txt := s.Text()
			txt = strings.TrimSpace(txt)
			re := regexp.MustCompile(`(\[)+\d+(\])`)
			txtClean1 := re.ReplaceAllString(txt, "")

			atricles = append(atricles, txtClean1)
		}
		i++
	})

	return atricles, nil

}
