package api

import (
	"fmt"
	"io"
	"net/url"
)

type PatchOptions struct {
	ContentType string
}

func NewPatchOptions() *PatchOptions {
	return &PatchOptions{
		ContentType: "application/json",
	}
}

func Patch(resource string, Req io.ReadCloser, options *PatchOptions) (body io.ReadCloser, err error) {
	Url := fmt.Sprintf("%v/%v", _host, resource)
	u, err := url.Parse(Url)
	if err != nil {
		return
	}
	return Request("PATCH", u, Req, options.ContentType)
}

// HelpPatchText HelpPatchText
var HelpPatchText = `
For most of Operations in FHIR you manipulate a resource as a whole (create, update, delete operations). But sometimes you want to update specific data elements in a resource and do not care about the rest. In other words, you need an element/attribute level operation.
`
