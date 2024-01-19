package httprouter

import (
	"net/http"
	"strconv"
)

// PaginationRequest represents a request for pagination.
type PaginationRequest struct {
	page    int // The current page number.
	perPage int // The number of items per page.
}

// ParsePaginationRequest parses the pagination parameters from the given HTTP request
// and returns a PaginationRequest object.
// The `limit` parameter specifies the maximum number of items to be returned per page.
// The `offset` parameter specifies the number of items to skip before starting to return items.
// The `page` parameter specifies the current page number.
// The `per_page` parameter specifies the number of items to be returned per page,
// if not explicitly specified, it defaults to the limit value.
func ParsePaginationRequest(r *http.Request, defaultLimit int) PaginationRequest {
	// Set the default limit.
	if defaultLimit == 0 {
		defaultLimit = 10
	}

	// Set the default pagination request.
	p := PaginationRequest{
		page:    1,
		perPage: defaultLimit,
	}

	// Return the default pagination request if the HTTP request is nil.
	if r == nil {
		return p
	}

	// Get the page.
	if page := GetQueryInt(r, "page"); page > 0 {
		p.page = page
	}

	// Get the per page.
	if perPage := GetQueryInt(r, "per_page"); perPage > 0 {
		p.perPage = perPage
	}

	// Return the pagination request.
	return p
}

// Limit returns the maximum number of items to return.
func (r PaginationRequest) Limit() int {
	return r.PerPage()
}

// Offset returns the number of items to skip.
func (r PaginationRequest) Offset() int {
	return (r.Page() - 1) * r.PerPage()
}

// Page returns the current page number.
func (r PaginationRequest) Page() int {
	return r.page
}

// PerPage returns the number of items per page.
func (r PaginationRequest) PerPage() int {
	return r.perPage
}

// GetQueryInt returns the value of the given query parameter as an integer.
// If the query parameter is not found, it returns 0.
func GetQueryInt(r *http.Request, key string) int {
	// Get the query parameter.
	value := r.URL.Query().Get(key)
	if value == "" {
		// Return the default value if the query parameter is not found.
		return 0
	}

	// Parse the query parameter.
	if intValue, err := strconv.Atoi(value); err == nil {
		return intValue
	}

	// Return the default value.
	return 0
}
