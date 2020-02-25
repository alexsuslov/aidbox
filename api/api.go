package api

import (
	"io"
	"log"
	"os"
)

var (
	// DEBUGGING DEBUGGING
	DEBUGGING bool
	// _InsecureSkipVerify bool
	// _host               string
	// _client             string
	// _secret             string
)

// Client Client
type Client struct {
	host               string
	client             string
	secret             string
	insecureSkipVerify bool
}

// New New
func New(host string, client1 string, secret string, insecureSkipVerify bool) (client *Client, err error) {
	client = &Client{
		host:               host,
		client:             client1,
		secret:             secret,
		insecureSkipVerify: insecureSkipVerify,
	}
	// TODO: check health aidbox
	return
}

// Print Print
func Print(opts ...interface{}) {
	if DEBUGGING {
		log.Println(opts...)
	}
}

var contentType = map[string]string{
	"json": "application/json",
	"xml":  "text/xml",
	"yml":  "text/yaml",
	"yaml": "text/yaml",
}

// Done Done
func Done(body io.ReadCloser, err error) {
	if err != nil {
		panic(err)
	}
	defer body.Close()
	if _, err := io.Copy(os.Stdout, body); err != nil {
		panic(err)

	}
}
