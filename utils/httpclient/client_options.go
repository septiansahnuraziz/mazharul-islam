package httpclient

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"go.opentelemetry.io/otel/attribute"
)

// Transport for tracing HTTP operations.
type Transport struct {
	rt http.RoundTripper
}

// Option signature for specifying options, e.g. WithRoundTripper.
type Option func(t *Transport)

// WithRoundTripper specifies the http.RoundTripper to call
// next after this transport. If it is nil (default), the
// transport will use http.DefaultTransport.
func WithRoundTripper(rt http.RoundTripper) Option {
	return func(t *Transport) {
		t.rt = rt
	}
}

// NewTransport specifies a transport that will trace HTTP
// and report back via OpenTracing.
func NewTransport(opts ...Option) *Transport {
	t := &Transport{}
	for _, o := range opts {
		o(t)
	}
	return t
}

// RoundTrip captures the request and starts an OpenTracing span
// for HTTP operation.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {

	// See General (https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/semantic_conventions/span-general.md)
	// and HTTP (https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/semantic_conventions/http.md)
	attributes := []attribute.KeyValue{
		attribute.String("http.url", req.URL.Redacted()),
		attribute.String("http.method", req.Method),
		attribute.String("http.scheme", req.URL.Scheme),
		attribute.String("http.host", req.URL.Hostname()),
		attribute.String("http.path", req.URL.Path),
		attribute.String("http.user_agent", req.UserAgent()),
	}

	var (
		buf    []byte
		err    error
		reader io.ReadCloser
	)
	if req.Body == nil {
		goto SetAttribute
	}

	buf, err = ioutil.ReadAll(req.Body)
	if err == nil {
		attributes = append(attributes, attribute.String("http.body", string(buf)))
	}

	reader = ioutil.NopCloser(bytes.NewBuffer(buf))
	req.Body = reader
SetAttribute:

	var (
		resp *http.Response
	)

	resp, err = t.rt.RoundTrip(req)
	if t.rt == nil {
		resp, err = http.DefaultTransport.RoundTrip(req)
	}

	if err != nil {
		log.Println("http client round trip error:", err)
	}

	if resp != nil {
		logrus.Error(err)
	}

	return resp, err
}
