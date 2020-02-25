package api

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"os"
)

// PatchOptions PatchOptions
type PatchOptions struct {
	ContentType string
}

// NewPatchOptions NewPatchOptions
func NewPatchOptions() *PatchOptions {
	return &PatchOptions{
		ContentType: "application/json",
	}
}

// Patch Patch
func (Client Client) Patch(resource string, Req io.ReadCloser,
	options *PatchOptions) (body io.ReadCloser, err error) {
	URL := fmt.Sprintf("%v/%v", Client.host, resource)
	u, err := url.Parse(URL)
	if err != nil {
		return
	}
	return Client.Request("PATCH", u, Req, options.ContentType)
}

// HelpPatchText Help Patch Text
var HelpPatchText = `
For most of Operations in FHIR you manipulate a resource as a whole (create, update, delete operations). But sometimes you want to update specific data elements in a resource and do not care about the rest. In other words, you need an element/attribute level operation.
`

// MainPatch MainPatch
func (Client Client) MainPatch(ctype string, patch string) {
	if patch != "" {
		t, ok := contentType[ctype]
		if !ok {
			panic("Error Content-Type")
		}
		reader := bufio.NewReader(os.Stdin)
		body, err := Client.Patch(patch, ioutil.NopCloser(reader), &PatchOptions{ContentType: t})
		Done(body, err)
		os.Exit(0)
	}
}
