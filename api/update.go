package api

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"os"
)

// UpdateOptions UpdateOptions
type UpdateOptions struct {
	ContentType string
}

// NewUpdateOptions NewUpdateOptions
func NewUpdateOptions() *UpdateOptions {
	return &UpdateOptions{
		ContentType: "application/json",
	}
}

// Update Update
func (Client Client) Update(resource string, Req io.ReadCloser, options *UpdateOptions) (body io.ReadCloser, err error) {
	URL := fmt.Sprintf("%v/%v", Client.host, resource)
	u, err := url.Parse(URL)
	if err != nil {
		return
	}
	return Client.Request("PUT", u, Req, options.ContentType)
}

// MainUpdate MainUpdate
func (Client Client) MainUpdate(ctype string, update string) {
	if update != "" {
		t, ok := contentType[ctype]
		if !ok {
			panic("Error Content-Type")
		}
		reader := bufio.NewReader(os.Stdin)
		body, err := Client.Update(update, ioutil.NopCloser(reader), &UpdateOptions{ContentType: t})
		Done(body, err)
		os.Exit(0)
	}
}
