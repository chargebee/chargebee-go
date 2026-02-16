package testutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"

	chargebee "github.com/chargebee/chargebee-go/v4"
)

const (
	DefaultSiteName = "test-site"
	DefaultAPIKey   = "test-api-key"
)

// RequestAssertion lets tests validate an incoming request.
type RequestAssertion func(req *http.Request, body []byte) error

// MockResponse represents one queued response.
type MockResponse struct {
	StatusCode int
	Body       []byte
	Headers    http.Header
	Assert     RequestAssertion
}

// CapturedRequest stores data from each incoming request.
type CapturedRequest struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// Server is an HTTP test server that queues responses and records requests.
type Server struct {
	mu               sync.Mutex
	server           *httptest.Server
	targetURL        *url.URL
	queue            []MockResponse
	requests         []CapturedRequest
	assertionFailure error
}

// NewServer starts a new mock server instance.
func NewServer() *Server {
	s := &Server{}
	s.server = httptest.NewServer(http.HandlerFunc(s.handle))
	s.targetURL, _ = url.Parse(s.server.URL)
	return s
}

// Close shuts down the underlying test server.
func (s *Server) Close() {
	if s == nil || s.server == nil {
		return
	}
	s.server.Close()
}

// URL returns the underlying test server URL.
func (s *Server) URL() string {
	if s == nil || s.server == nil {
		return ""
	}
	return s.server.URL
}

// Enqueue adds a response to be returned for the next request.
func (s *Server) Enqueue(res MockResponse) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if res.StatusCode == 0 {
		res.StatusCode = http.StatusOK
	}
	s.queue = append(s.queue, res)
}

// EnqueueJSON marshals payload as JSON and queues it as the next response.
func (s *Server) EnqueueJSON(statusCode int, payload any) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal mock payload: %w", err)
	}
	s.Enqueue(MockResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers: http.Header{
			"Content-Type": []string{"application/json"},
		},
	})
	return nil
}

// ClientConfig returns a Chargebee config wired to use this mock server.
func (s *Server) ClientConfig() *chargebee.ClientConfig {
	return &chargebee.ClientConfig{
		ApiKey:   DefaultAPIKey,
		SiteName: DefaultSiteName,
		HTTPClient: &http.Client{
			Transport: &rewriteTransport{
				targetURL: s.targetURL,
				base:      http.DefaultTransport,
			},
		},
	}
}

// NewClient creates a Chargebee client pre-configured to hit this server.
func (s *Server) NewClient() *chargebee.Client {
	return chargebee.NewClient(s.ClientConfig())
}

// Requests returns a snapshot of all captured requests.
func (s *Server) Requests() []CapturedRequest {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]CapturedRequest, len(s.requests))
	copy(out, s.requests)
	return out
}

// LastRequest returns the most recent captured request.
func (s *Server) LastRequest() (CapturedRequest, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.requests) == 0 {
		return CapturedRequest{}, false
	}
	return s.requests[len(s.requests)-1], true
}

// PendingResponses returns how many queued responses are left.
func (s *Server) PendingResponses() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.queue)
}

// AssertionError returns the first assertion error that occurred.
func (s *Server) AssertionError() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.assertionFailure
}

func (s *Server) handle(w http.ResponseWriter, req *http.Request) {
	body, _ := io.ReadAll(req.Body)
	req.Body.Close()

	s.mu.Lock()
	captured := CapturedRequest{
		Method: req.Method,
		Path:   req.URL.Path,
		Query:  req.URL.Query(),
		Header: req.Header.Clone(),
		Body:   bytes.Clone(body),
	}
	s.requests = append(s.requests, captured)

	if len(s.queue) == 0 {
		s.mu.Unlock()
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error":"no mock response queued"}`))
		return
	}

	mockResponse := s.queue[0]
	s.queue = s.queue[1:]
	s.mu.Unlock()

	if mockResponse.Assert != nil {
		if err := mockResponse.Assert(req, body); err != nil {
			s.mu.Lock()
			if s.assertionFailure == nil {
				s.assertionFailure = err
			}
			s.mu.Unlock()
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
	}

	for key, values := range mockResponse.Headers {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(mockResponse.StatusCode)
	if len(mockResponse.Body) > 0 {
		_, _ = w.Write(mockResponse.Body)
	}
}

type rewriteTransport struct {
	targetURL *url.URL
	base      http.RoundTripper
}

func (t *rewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	base := t.base
	if base == nil {
		base = http.DefaultTransport
	}

	clone := req.Clone(req.Context())
	clone.URL.Scheme = t.targetURL.Scheme
	clone.URL.Host = t.targetURL.Host
	clone.RequestURI = ""
	if clone.Host == "" || strings.Contains(clone.Host, "chargebee.com") {
		clone.Host = t.targetURL.Host
	}

	return base.RoundTrip(clone)
}
