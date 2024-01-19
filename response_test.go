package httprouter_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dmitrymomot/httprouter"
	"github.com/stretchr/testify/require"
)

func TestJSON(t *testing.T) {
	t.Parallel()

	// Test case 1: result is nil
	t.Run("result is nil", func(t *testing.T) {
		w := httptest.NewRecorder()
		require.NoError(t, httprouter.JSON(w, http.StatusOK, nil))
		require.Equal(t, http.StatusNoContent, w.Code)
		require.Equal(t, "", w.Body.String())
	})

	// Test case 2: result is not nil
	t.Run("result is not nil", func(t *testing.T) {
		expectedResult := map[string]interface{}{
			"message": "Hello, World!",
		}

		w := httptest.NewRecorder()
		require.NoError(t, httprouter.JSON(w, http.StatusOK, expectedResult))
		require.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		require.NoError(t, json.NewDecoder(w.Body).Decode(&response))
		require.Equal(t, expectedResult, response)
	})

	// Test case 3: failed to encode response
	t.Run("failed to encode response", func(t *testing.T) {
		w := httptest.NewRecorder()
		require.Error(t, httprouter.JSON(w, http.StatusOK, make(chan int)))
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, "", w.Body.String())
	})
}

func TestString(t *testing.T) {
	t.Parallel()

	// Test case 1: result is nil
	t.Run("result is nil", func(t *testing.T) {
		w := httptest.NewRecorder()
		require.NoError(t, httprouter.PlainText(w, http.StatusOK, ""))
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, "", w.Body.String())
	})

	// Test case 2: result is not nil
	t.Run("result is not nil", func(t *testing.T) {
		expectedResult := "Hello, World!"

		w := httptest.NewRecorder()
		require.NoError(t, httprouter.PlainText(w, http.StatusOK, expectedResult))
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, expectedResult, w.Body.String())
	})
}

func TestHTML(t *testing.T) {
	t.Parallel()

	// Test case 1: result is nil
	t.Run("result is nil", func(t *testing.T) {
		w := httptest.NewRecorder()
		require.ErrorIs(t, httprouter.HTML(w, http.StatusOK, []byte(nil)), httprouter.ErrResponsePayloadIsNil)
		require.Equal(t, "", w.Body.String())
	})

	// Test case 2: result is not nil
	t.Run("result is not nil", func(t *testing.T) {
		expectedResult := "<html><body>Hello, World!</body></html>"

		w := httptest.NewRecorder()
		require.NoError(t, httprouter.HTML(w, http.StatusOK, []byte(expectedResult)))
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, expectedResult, w.Body.String())
	})
}

func TestNoContent(t *testing.T) {
	t.Parallel()

	// Test case 1: success
	t.Run("success", func(t *testing.T) {
		w := httptest.NewRecorder()
		require.NoError(t, httprouter.NoContent(w))
		require.Equal(t, http.StatusNoContent, w.Code)
		require.Equal(t, "", w.Body.String())
	})
}

func TestRedirect(t *testing.T) {
	t.Parallel()

	// Test case 1: success
	t.Run("success", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/redirect", nil)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		err = httprouter.Redirect(w, req, http.StatusFound, "https://example.com")
		require.NoError(t, err)

		require.Equal(t, http.StatusFound, w.Code)
		require.Equal(t, "https://example.com", w.Header().Get("Location"))
	})
}
