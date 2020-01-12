package api

import (
	"fmt"
	"io"
	"net/url"
)

type CreateOptions struct {
	ContentType string
}

func NewCreateOptions() *CreateOptions{
	return &CreateOptions{
		ContentType: "application/json",
	}
}

func Create(resource string, Req io.ReadCloser, options *CreateOptions, )(body io.ReadCloser, err error) {
	Url := fmt.Sprintf("%v/%v", _host, resource)
	u, err := url.Parse(Url)
	if err!= nil{
		return
	}
	return  Request("POST", u, Req, options.ContentType)
}
