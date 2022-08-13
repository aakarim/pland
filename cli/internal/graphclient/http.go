package graphclient

import "net/http"

type AuthedTransport struct {
	key     string
	wrapped http.RoundTripper
}

func NewAuthedTransport(key string, wrapped http.RoundTripper) *AuthedTransport {
	return &AuthedTransport{
		key:     key,
		wrapped: wrapped,
	}
}

func (t *AuthedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "bearer "+t.key)
	resp, err := t.wrapped.RoundTrip(req)
	return resp, err
}
