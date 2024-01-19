package httprouter

import (
	"encoding/json"
	"errors"
	"net/http"
)

// JSON is a response encoder that encodes the given result as JSON and writes it to the http.ResponseWriter.
// It sets the Content-Type header to "application/json; charset=utf-8" and the status code to the provided code.
// If the result is nil, it returns a response with status code http.StatusNoContent.
func JSON(w http.ResponseWriter, code int, result interface{}) error {
	// Check if the response is not nil.
	if result == nil {
		// If the result is nil, return a response with status code http.StatusNoContent.
		w.WriteHeader(http.StatusNoContent)
		return nil
	}

	// Encode the response.
	w.Header().Set(ContentTypeHeader, ContentTypeJSONWithCharset)
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		return errors.Join(ErrFailedToEncodeResponse, err)
	}

	return nil
}

// PlainText writes a plain text response to the http.ResponseWriter.
// It sets the provided HTTP status code and encodes the response as plain text.
// Returns an error if there was a failure in encoding the response.
func PlainText[Payload string | []byte](w http.ResponseWriter, code int, payload Payload) error {
	// Encode the response.
	w.Header().Set(ContentTypeHeader, ContentTypePlainTextWithCharset)
	w.WriteHeader(code)
	if _, err := w.Write([]byte(payload)); err != nil {
		return errors.Join(ErrFailedToEncodeResponse, err)
	}

	return nil
}

// HTML encodes the given result as HTML and writes it to the http.ResponseWriter.
// It sets the Content-Type header to "text/html; charset=utf-8" and the HTTP status code to the provided code.
// If the result is nil, it returns an error indicating that the response cannot be encoded.
// If the encoding or writing of the response fails, it returns an error.
func HTML[Payload string | []byte](w http.ResponseWriter, code int, payload Payload) error {
	// Check if the response is not nil.
	if []byte(payload) == nil {
		// Return an error if the response is nil.
		// This is because the HTML encoder cannot encode nil responses.
		// If you want to return a nil response, use the NoContent encoder.
		return errors.Join(ErrFailedToEncodeResponse, ErrResponsePayloadIsNil)
	}

	// Encode the response.
	w.Header().Set(ContentTypeHeader, ContentTypeHTMLWithCharset)
	w.WriteHeader(code)
	if _, err := w.Write([]byte(payload)); err != nil {
		return errors.Join(ErrFailedToEncodeResponse, err)
	}

	return nil
}

// NoContent writes a response with HTTP status code 204 (No Content).
// It sets the status code in the provided http.ResponseWriter to 204 and returns nil.
func NoContent(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusNoContent)
	return nil
}

// Redirect redirects the HTTP request to the specified URL with the given status code.
// It sets the Location header in the response to the provided URL and sends the redirect response to the client.
// The function takes the http.ResponseWriter, *http.Request, status code, and URL as parameters.
// It returns an error if there was an issue redirecting the request.
func Redirect(w http.ResponseWriter, r *http.Request, code int, url string) error {
	http.Redirect(w, r, url, code)
	return nil
}
