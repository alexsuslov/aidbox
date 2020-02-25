package api

import (
	"fmt"
	"io"
	"net/url"
)

// DeleteOptions DeleteOptions
type DeleteOptions struct {
	ContentType string
}

// NewDeleteOptions NewDeleteOptions
func (Client Client) NewDeleteOptions() *DeleteOptions {
	return &DeleteOptions{
		ContentType: "application/json",
	}
}

// Delete Delete
func (Client Client) Delete(resource string, options *DeleteOptions) (body io.ReadCloser, err error) {
	URL := fmt.Sprintf("%v/%v", Client.host, resource)
	u, err := url.Parse(URL)
	if err != nil {
		return
	}
	return Client.Request("DELETE", u, nil, options.ContentType)
}
