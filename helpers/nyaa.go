package helpers

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type NyaaSortBy string

const (
	NyaaSortBySize      NyaaSortBy = "size"
	NyaaSortByDate      NyaaSortBy = "id"
	NyaaSortBySeeders   NyaaSortBy = "seeders"
	NyaaSortByLeechers  NyaaSortBy = "leechers"
	NyaaSortByDownloads NyaaSortBy = "downloads"
)

var NyaaSortByMap = map[NyaaSortBy]string{
	NyaaSortBySize:      "Size",
	NyaaSortByDate:      "Date",
	NyaaSortBySeeders:   "Seeders",
	NyaaSortByLeechers:  "Leechers",
	NyaaSortByDownloads: "Downloads",
}

func (x NyaaSortBy) Title() string {
	return NyaaSortByMap[x]
}

func (x NyaaSortBy) IsValid() bool {
	_, ok := NyaaSortByMap[x]
	return ok
}

type NyaaSortOrder string

const (
	NyaaSortOrderAscending  NyaaSortOrder = "asc"
	NyaaSortOrderDescending NyaaSortOrder = "desc"
)

var NyaaSortOrderMap = map[NyaaSortOrder]string{
	NyaaSortOrderAscending:  "Ascending",
	NyaaSortOrderDescending: "Descending",
}

func (x NyaaSortOrder) Title() string {
	return NyaaSortOrderMap[x]
}

func (x NyaaSortOrder) IsValid() bool {
	_, ok := NyaaSortOrderMap[x]
	return ok
}

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

func NyaaSearch(terms string, page int, sort NyaaSortBy, order NyaaSortOrder) ([]NyaaSearchItem, error) {
	if terms == "" || page < 1 || !sort.IsValid() || !order.IsValid() {
		return []NyaaSearchItem{}, nil
	}
	base := "https://nyaa.si"
	url := fmt.Sprintf("%s/?q=%s&s=%s&o=%s&p=%d", base, url.QueryEscape(terms), sort, order, page)
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
