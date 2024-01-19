package httprouter

import (
	"net/http"
	"strings"
)

const (
	// ContentTypeHeader is the name of the Content-Type header.
	ContentTypeHeader = "Content-Type"
	// AcceptContentTypeHeader is the name of the Accept header.
	AcceptContentTypeHeader = "Accept"

	// ContentTypeCharsetUTF8 is the value of the charset parameter for the Content-Type header.
	ContentTypeCharsetUTF8 = "charset=utf-8"

	// Predefined content types.
	ContentTypeJSON      = "application/json"
	ContentTypeHTML      = "text/html"
	ContentTypePlainText = "text/plain"

	// Predefined content types with charset.
	ContentTypeJSONWithCharset      = "application/json; " + ContentTypeCharsetUTF8
	ContentTypeHTMLWithCharset      = "text/html; " + ContentTypeCharsetUTF8
	ContentTypePlainTextWithCharset = "text/plain; " + ContentTypeCharsetUTF8
)

// IsJSONRequest returns true if the given request has a Content-Type header with the value "application/json".
func IsJSONRequest(r *http.Request) bool {
	return IsContentType(r, ContentTypeJSON)
}

// IsHTMLRequest returns true if the given request has a Content-Type header with the value "text/html".
func IsHTMLRequest(r *http.Request) bool {
	return IsContentType(r, ContentTypeHTML)
}

// IsPlainTextRequest returns true if the given request has a Content-Type header with the value "text/plain".
func IsPlainTextRequest(r *http.Request) bool {
	return IsContentType(r, ContentTypePlainText)
}

// Split the Content-Type header into its media type and parameters.
// The media type is the first part of the header value.
// The parameters are the second part of the header value.
// For example, the Content-Type header "application/json; charset=utf-8" has the media type "application/json" and the parameter "charset=utf-8".
func SplitContentTypeHeader(header string) (mediaType string, parameters string) {
	// Convert the header to lowercase and trim whitespace.
	header = strings.ToLower(strings.TrimSpace(header))

	// Split the header into its parts.
	parts := strings.SplitN(header, ";", 2)

	// Set the media type.
	mediaType = parts[0]

	// Set the parameters.
	if len(parts) > 1 {
		parameters = parts[1]
	}

	return mediaType, parameters
}

// IsContentType returns true if the given request has a Content-Type header with the given value.
func IsContentType(r *http.Request, contentType string) bool {
	// Get the Content-Type header.
	header := r.Header.Get(ContentTypeHeader)

	// Split the Content-Type header into its media type and parameters.
	mediaType, _ := SplitContentTypeHeader(header)

	// Compare the media type with the given value.
	return mediaType == contentType
}
