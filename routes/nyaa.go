package routes

import (
	"net/http"
	"sailing-seas/core"
	"sailing-seas/helpers"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

func NyaaRoute(app *core.App, mux *http.ServeMux) {
	mux.HandleFunc("GET /nyaa", nyaaRouteHandler())
}

func nyaaRouteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		x := newNyaaSearch()
		x.Apply(r)
		x.Execute()
		if x.Err != nil {
			log.Error().Err(x.Err).Msgf("nyaa search failed for query '%s'", x.Terms)
		}
		queries := r.URL.Query()
		queries.Del("p")
		queryNoPage := queries.Encode()
		err := ExecuteTemplate(w, "nyaa", map[string]any{
			"SortByMap":      helpers.NyaaSortByMap,
			"SortOrderMap":   helpers.NyaaSortOrderMap,
			"SearchTerms":    x.Terms,
			"Page":           x.Page,
			"SortBy":         x.Sort,
			"SortOrder":      x.Order,
			"SearchResults":  x.Items,
			"UrlQueryNoPage": queryNoPage,
			"ErrorMessage":   x.Err,
		})
		if err != nil {
			log.Error().Err(err).Msg("failed to render /nyaa")
		}
	}
}

type nyaaSearch struct {
	Terms string
	Page  int
	Sort  helpers.NyaaSortBy
	Order helpers.NyaaSortOrder
	Items []helpers.NyaaSearchItem
	Err   error
}

func newNyaaSearch() *nyaaSearch {
	return &nyaaSearch{
		Terms: "",
		Page:  1,
		Sort:  helpers.NyaaSortBySeeders,
		Order: helpers.NyaaSortOrderDescending,
		Items: []helpers.NyaaSearchItem{},
		Err:   nil,
	}
}

func (x *nyaaSearch) Apply(r *http.Request) {
	queries := r.URL.Query()
	x.Terms = strings.TrimSpace(queries.Get("q"))
	if pageStr := queries.Get("p"); pageStr != "" {
		x.Page, x.Err = strconv.Atoi(pageStr)
	}
	if sortStr := queries.Get("s"); sortStr != "" {
		x.Sort = helpers.NyaaSortBy(sortStr)
	}
	if orderStr := queries.Get("o"); orderStr != "" {
		x.Order = helpers.NyaaSortOrder(orderStr)
	}
}

func (x *nyaaSearch) Execute() {
	x.Items, x.Err = helpers.NyaaSearch(x.Terms, x.Page, x.Sort, x.Order)
}
