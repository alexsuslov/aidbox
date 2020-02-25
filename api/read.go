package api

import (
	"fmt"
	"io"
	"net/url"
	"os"
)

// ReadOptions ReadOptions
type ReadOptions struct {
	ContentType string
}

// NewReadOptions NewReadOptions
func NewReadOptions() *ReadOptions {
	return &ReadOptions{
		ContentType: "application/json",
	}
}

// Read Read
func (Client Client) Read(resource string, options *ReadOptions) (body io.ReadCloser, err error) {
	URL := fmt.Sprintf("%v/%v", Client.host, resource)
	u, err := url.Parse(URL)
	if err != nil {
		return
	}
	return Client.Request("GET", u, nil, options.ContentType)
}

// MainRead MainRead
func (Client Client) MainRead(ctype string, read string) {
	if read != "" {
		t, ok := contentType[ctype]
		if !ok {
			panic("Error Content-Type")
		}
		body, err := Client.Read(read, &ReadOptions{ContentType: t})
		Done(body, err)
		os.Exit(0)
	}
}
