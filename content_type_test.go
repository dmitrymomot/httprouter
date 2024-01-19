package httprouter_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dmitrymomot/httprouter"
	"github.com/stretchr/testify/require"
)

func TestIsJSONRequest(t *testing.T) {
	t.Parallel()

	// Test case 1: JSON content type
	t.Run("JSON content type", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Content-Type", "application/json")

		result := httprouter.IsJSONRequest(req)
		require.True(t, result)
	})

	// Test case 2: non-JSON content type
	t.Run("non-JSON content type", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Content-Type", "text/plain")

		result := httprouter.IsJSONRequest(req)
		require.False(t, result)
	})

	// Test case 3: missing content type header
	t.Run("missing content type header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)

		result := httprouter.IsJSONRequest(req)
		require.False(t, result)
	})
}

func TestIsHTMLRequest(t *testing.T) {
	t.Parallel()

	// Test case 1: HTML content type
	t.Run("HTML content type", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Content-Type", "text/html")

		result := httprouter.IsHTMLRequest(req)
		require.True(t, result)
	})

	// Test case 2: non-HTML content type
	t.Run("non-HTML content type", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Content-Type", "application/json")

		result := httprouter.IsHTMLRequest(req)
		require.False(t, result)
	})

	// Test case 3: missing content type header
	t.Run("missing content type header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)

		result := httprouter.IsHTMLRequest(req)
		require.False(t, result)
	})
}

func TestIsPlainTextRequest(t *testing.T) {
	t.Parallel()

	// Test case 1: plain text content type
	t.Run("plain text content type", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Content-Type", "text/plain")

		result := httprouter.IsPlainTextRequest(req)
		require.True(t, result)
	})

	// Test case 2: non-plain text content type
	t.Run("non-plain text content type", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Content-Type", "application/json")

		result := httprouter.IsPlainTextRequest(req)
		require.False(t, result)
	})

	// Test case 3: missing content type header
	t.Run("missing content type header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)

		result := httprouter.IsPlainTextRequest(req)
		require.False(t, result)
	})
}
