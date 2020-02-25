package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"

	"gopkg.in/yaml.v2"
)

// EntityOptions EntityOptions
type EntityOptions struct {
	ContentType string
}

// NewEntityOptions NewEntityOptions
func NewEntityOptions() *EntityOptions {
	return &EntityOptions{
		ContentType: "application/json",
	}
}

// Entity Entity
func (Client Client) Entity(Req io.ReadCloser,
	options *EntityOptions) (body io.ReadCloser, err error) {
	values := url.Values{}
	if err = parseEntityOptions(options.ContentType, Req, &values); err != nil {
		return
	}

	URL := fmt.Sprintf("%v/%v", Client.host, values.Encode())
	u, err := url.Parse(URL)
	if err != nil {
		return
	}
	return Client.Request("GET", u, Req, options.ContentType)
}

func parseEntityOptions(ReqType string, Req io.ReadCloser, values *url.Values) (err error) {
	Parsed := map[string]string{}
	switch ReqType {
	case "application/json":
		err = json.NewDecoder(Req).Decode(&Parsed)
		if err != nil {
			return err
		}
	case "application/yaml":
		err = yaml.NewDecoder(Req).Decode(&Parsed)
		if err != nil {
			return err
		}
	}
	for name, value := range Parsed {
		values.Add(name, value)
	}
	return nil
}

// MainEntity MainEntity
func (Client Client) MainEntity(entity bool, ctype string, req io.ReadCloser) {
	if entity {
		t, ok := contentType[ctype]
		if !ok {
			panic("Error Content-Type")
		}
		body, err := Client.Entity(req, &EntityOptions{
			ContentType: t,
		})
		Done(body, err)
		os.Exit(0)
	}
}

// HelperEntity HelperEntity
func HelperEntity() {

}
