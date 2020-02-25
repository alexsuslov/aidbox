package api

import (
	"fmt"
	"io"
	"net/url"
)

// CreateOptions CreateOptions
type CreateOptions struct {
	ContentType string
}

// NewCreateOptions NewCreateOptions
func NewCreateOptions() *CreateOptions {
	return &CreateOptions{
		ContentType: "application/json",
	}
}

// Create Create
func (Client Client) Create(resource string, Req io.ReadCloser, options *CreateOptions) (body io.ReadCloser, err error) {
	URL := fmt.Sprintf("%v/%v", Client.host, resource)
	u, err := url.Parse(URL)
	if err != nil {
		return
	}
	return Client.Request("POST", u, Req, options.ContentType)
}
