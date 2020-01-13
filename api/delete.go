package api

import (
	"fmt"
	"io"
	"net/url"
)

type DeleteOptions struct {
	ContentType string
}

func NewDeleteOptions() *DeleteOptions{
	return &DeleteOptions{
		ContentType: "application/json",
	}
}

func Delete(resource string, options *DeleteOptions)(body io.ReadCloser, err error) {
	Url := fmt.Sprintf("%v/%v", _host, resource)
	u, err := url.Parse(Url)
	if err!= nil{
		return
	}
	return  Request("DELETE", u, nil, options.ContentType)
}
