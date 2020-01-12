package api

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Request Request
func Request(method string, url *url.URL, reader io.ReadCloser, contentType string) (body io.ReadCloser, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: _InsecureSkipVerify},
	}
	Url := url.String()
	Print("method", method)
	Print("url", Url)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	req, err := http.NewRequestWithContext(ctx, method, Url, reader)
	if err!= nil{
		return
	}
	basic := base64.StdEncoding.EncodeToString(
		[]byte(
			fmt.Sprintf("%v:%v", _client, _secret),
			),
		)
	authorization := fmt.Sprintf("Basic %v", basic)
	req.Header.Set("Authorization", authorization)
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{Transport: tr}
	r, err := client.Do(req)
	if err != nil {
		Print("error client.Do:", err)
		err = fmt.Errorf("client.Do:%v", err)
		return
	}
	if r.StatusCode == 500 {
		data, _ := ioutil.ReadAll(r.Body)
		err = errors.New(string(data))
		Print("error server side:", r.StatusCode, err)
		err = fmt.Errorf("StatusCode:%v", r.StatusCode)
		return
	}

	if r.StatusCode < 200 || r.StatusCode >= 300 {
		err = fmt.Errorf("status:%v", r.Status)
		return
	}

	return r.Body, err
}