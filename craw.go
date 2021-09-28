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
func Fetch(url string) ([]string, error) {
	var atricles []string
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

	rawTxt := rawArticles.Text()

	splitted := strings.Split(rawTxt, ".")

	for i, v := range splitted {

		isSmall := len(strings.Join(atricles, "")) < 2500

		if i < ARTICLES_TO_TAKE || isSmall {
			txt := v
			txt = strings.TrimSpace(txt)
			re := regexp.MustCompile(`(\[)+\d+(\])`)
			txtClean1 := re.ReplaceAllString(txt, "") + "."

			atricles = append(atricles, txtClean1)
		}
	}

	return atricles, nil

}
