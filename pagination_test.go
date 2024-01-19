package httprouter_test

import (
	"net/http"
	"testing"

	"github.com/dmitrymomot/httprouter"
	"github.com/stretchr/testify/require"
)

func TestGetQueryInt(t *testing.T) {
	// Test case 1: query parameter is not found, return default value
	t.Run("query parameter not found", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/path", nil)
		require.NoError(t, err)

		result := httprouter.GetQueryInt(req, "page")
		require.Equal(t, 0, result)
	})

	// Test case 2: query parameter is found and can be parsed as int
	t.Run("query parameter found and parsed successfully", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/path?page=5", nil)
		require.NoError(t, err)

		result := httprouter.GetQueryInt(req, "page")
		require.Equal(t, 5, result)
	})

	// Test case 3: query parameter is found but cannot be parsed as int, return default value
	t.Run("query parameter found but cannot be parsed as int", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/path?page=invalid", nil)
		require.NoError(t, err)

		result := httprouter.GetQueryInt(req, "page")
		require.Equal(t, 0, result)
	})
}

func TestParsePaginationRequest(t *testing.T) {
	defaultLimit := 20

	// Test case 0: default limit is not set
	t.Run("default limit is not set", func(t *testing.T) {
		result := httprouter.ParsePaginationRequest(nil, 0)
		require.Equal(t, 10, result.Limit())
		require.Equal(t, 0, result.Offset())
		require.Equal(t, 1, result.Page())
		require.Equal(t, 10, result.PerPage())
	})

	// Test case 1: request is nil
	t.Run("request is nil", func(t *testing.T) {
		result := httprouter.ParsePaginationRequest(nil, defaultLimit)
		require.Equal(t, defaultLimit, result.Limit())
		require.Equal(t, 0, result.Offset())
		require.Equal(t, 1, result.Page())
		require.Equal(t, defaultLimit, result.PerPage())
	})

	// Test case 2: page is set
	t.Run("page is set", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/path?page=5", nil)
		require.NoError(t, err)

		result := httprouter.ParsePaginationRequest(req, defaultLimit)
		require.Equal(t, defaultLimit, result.Limit())
		require.Equal(t, 80, result.Offset())
		require.Equal(t, 5, result.Page())
		require.Equal(t, defaultLimit, result.PerPage())
	})

	// Test case 3: per page is set
	t.Run("per page is set", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/path?per_page=5", nil)
		require.NoError(t, err)

		result := httprouter.ParsePaginationRequest(req, defaultLimit)
		require.Equal(t, 5, result.Limit())
		require.Equal(t, 0, result.Offset())
		require.Equal(t, 1, result.Page())
		require.Equal(t, 5, result.PerPage())
	})

	// Test case 4: page and per page are set
	t.Run("page and per page are set", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/path?page=5&per_page=5", nil)
		require.NoError(t, err)

		result := httprouter.ParsePaginationRequest(req, defaultLimit)
		require.Equal(t, 5, result.Limit())
		require.Equal(t, 20, result.Offset())
		require.Equal(t, 5, result.Page())
		require.Equal(t, 5, result.PerPage())
	})

	// Test case 5: page and per page are set to invalid values
	t.Run("page and per page are set to invalid values", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/path?page=invalid&per_page=invalid", nil)
		require.NoError(t, err)

		result := httprouter.ParsePaginationRequest(req, defaultLimit)
		require.Equal(t, defaultLimit, result.Limit())
		require.Equal(t, 0, result.Offset())
		require.Equal(t, 1, result.Page())
		require.Equal(t, defaultLimit, result.PerPage())
	})
}
