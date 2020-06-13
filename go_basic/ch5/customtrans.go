package main

import (
	"net/http"
)

type OurCustomTransport struct {
	Transport http.RoundTripper
}

func (t *OurCustomTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}

	return http.DefaultTransport
}

func (t *OurCustomTransport) RoundTrip(req *http.Request) (http.Response, error) {
	// do something
	return t.transport().RoundTrip(req)
}

func (t *OurCustomTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func main() {
	t := &OurCustomTransport{
		// ...
	}

	c := t.Client()
	resp, err := c.Get("http://example.com")
}
