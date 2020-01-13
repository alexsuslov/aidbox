package api

import (
	"fmt"
	"io"
	"net/url"
)

type UpdateOptions struct {
	ContentType string
}

func NewUpdateOptions() *UpdateOptions{
	return &UpdateOptions{
		ContentType: "application/json",
	}
}

func Update(resource string, Req io.ReadCloser, options *UpdateOptions, )(body io.ReadCloser, err error) {
	Url := fmt.Sprintf("%v/%v", _host, resource)
	u, err := url.Parse(Url)
	if err!= nil{
		return
	}
	return  Request("PUT", u, Req, options.ContentType)
}
