package nks

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOrganizationsStub(t *testing.T) {
	fmt.Println("GetOrganizations testing")

	client := NewHTTPTestClient(func(req *http.Request) *http.Response {

		assert.Equal(t, "/orgs", req.URL.String(), "URL should be equal")
		assert.Equal(t, "GET", req.Method, "Method shoud be: GET")

		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`[]`)),
			Header:     make(http.Header),
		}
	})
	api := NewTestClient(client)
	body, err := api.GetOrganizations()

	assert.NoError(t, err)
	assert.Equal(t, []Organization{}, body)
}

func TestGetOrganizationStub(t *testing.T) {
	var mockID int = 0

	client := NewHTTPTestClient(func(req *http.Request) *http.Response {

		assert.Equal(t, fmt.Sprintf("%s%d", "/orgs/", mockID), req.URL.String(), "URL should be equal")
		assert.Equal(t, "GET", req.Method, "Method shoud be: GET")

		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`{}`)),
			Header:     make(http.Header),
		}
	})
	api := NewTestClient(client)
	body, err := api.GetOrganization(mockID)

	assert.NoError(t, err)
	assert.Equal(t, new(Organization), body)
}
