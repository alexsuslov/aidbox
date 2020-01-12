package api

import (
	"net/url"
	"path"
)

type ReadOptions struct {
	ContentType string
}

func NewReadOptions() *ReadOptions{
	return &ReadOptions{
		ContentType: "application/json",
	}
}

func Read(resource string, options *ReadOptions)(body io.ReadCloser, err error) {
	Url := path.Join(_host, resource)
	u, err := url.Parse(Url)
	if err!= nil{
		return
	}
	return  Request("GET", u, nil, options.ContentType)
}