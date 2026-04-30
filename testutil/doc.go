// Package testutil contains helpers for mocking Chargebee API requests in tests.
//
// Quick usage:
//
//	mock := testutil.NewServer()
//	defer mock.Close()
//
//	mock.Enqueue(testutil.MockResponse{
//		StatusCode: 200,
//		Body:       []byte(`{"customer":{"id":"cust_123"}}`),
//	})
//
//	client := client := mock.NewClient()
//	// Use client as usual. Requests are routed to the mock server.
//
// You can inspect incoming requests using Requests() / LastRequest(), and you can
// validate requests per response via MockResponse.Assert.
package testutil
