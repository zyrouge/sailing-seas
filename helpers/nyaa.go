package helpers

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type NyaaSearchItem struct {
	Category  string
	Title     string
	Url       string
	Magnet    string
	Size      string
	Date      string
	Seeders   int
	Leechers  int
	Downloads int
}

func NyaaSearch(terms string) ([]NyaaSearchItem, error) {
	if terms == "" {
		return []NyaaSearchItem{}, nil
	}
	base := "https://nyaa.si"
	url := fmt.Sprintf("%s/?q=%s&s=seeders&o=desc", base, url.QueryEscape(terms))
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	items := []NyaaSearchItem{}
	doc.Find(".torrent-list tbody tr").Each(func(i int, x *goquery.Selection) {
		rows := x.Find("td")
		name := rows.Eq(1).Find("a").Eq(-1)
		links := rows.Eq(2).Find("a")
		items = append(items, NyaaSearchItem{
			Category:  rows.Eq(0).Find("a").AttrOr("title", ""),
			Title:     name.Text(),
			Url:       base + name.AttrOr("href", ""),
			Magnet:    links.Eq(-1).AttrOr("href", ""),
			Size:      rows.Eq(3).Text(),
			Date:      rows.Eq(4).Text(),
			Seeders:   atoiOrZero(rows.Eq(5).Text()),
			Leechers:  atoiOrZero(rows.Eq(6).Text()),
			Downloads: atoiOrZero(rows.Eq(7).Text()),
		})
	})
	return items, nil
}

func atoiOrZero(value string) int {
	if parsed, err := strconv.Atoi(value); err == nil {
		return parsed
	}
	return 0
}
