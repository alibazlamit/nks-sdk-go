package nks

import (
	"net/http"
)

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

//NewHTTPTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewHTTPTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

//NewTestClient returns APIClient with Transport replaced to avoid making real calls
func NewTestClient(client *http.Client) *APIClient {
	c := &APIClient{
		Token:      "",
		Endpoint:   "",
		HttpClient: client,
	}
	return c
}
