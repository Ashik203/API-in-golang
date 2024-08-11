package utils

import (
	"math"
	"net/http"
	"strconv"
)

type PaginationParams struct {
	Page        int
	Limit       int
	SearchBy    string
	SearchValue string
	SortBy      string
	SortOrder   string
	FilterBy    string
	FilterValue string
}

const (
	maxLimit = 100.0
)

func parsePage(r *http.Request) int {
	pageStr := r.URL.Query().Get("pageKey")
	page, _ := strconv.ParseInt(pageStr, 10, 32)
	page = int64(math.Max(1.0, float64(page)))
	return int(page)
}

func parseLimit(r *http.Request) int {
	limitStr := r.URL.Query().Get("limitKey")
	limit, _ := strconv.ParseInt(limitStr, 10, 32)
	limit = int64(math.Max(0.0, math.Min(maxLimit, float64(limit))))
	return int(limit)
}

var defaultSortBy = "book_id"
var defaultSortOrder = "asc"

func GetPaginationParams(r *http.Request) PaginationParams {
	params := PaginationParams{
		Page:        1,
		Limit:       10,
		SearchBy:    "",
		SearchValue: "",
		SortBy:      defaultSortBy,
		SortOrder:   defaultSortOrder,
		FilterBy:    "",
		FilterValue: "",
	}

	queryParams := r.URL.Query()

	for k := range queryParams {
		switch k {
		case "pageKey":
			params.Page = parsePage(r)

		case "limitKey":
			params.Limit = parseLimit(r)

		case "searchKey":
			params.SearchBy = queryParams.Get("searchKey")

		case "searchValueKey":
			params.SearchValue = queryParams.Get("searchValueKey")

		case "sortByKey":
			params.SortBy = queryParams.Get("sortByKey")

		case "sortOrderKey":
			params.SortOrder = queryParams.Get("sortOrderKey")

		case "filterKey":
			params.FilterBy = queryParams.Get("filterKey")

		case "filterValueKey":
			params.FilterValue = queryParams.Get("filterValueKey")
		}
	}

	return params
}
