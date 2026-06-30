package telemetry

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildRequestHeaderSpanAttributesPromotesChargebeeHeaders(t *testing.T) {
	attributes := BuildRequestHeaderSpanAttributes(map[string]string{
		"chargebee-foo": "bar",
	})

	require.Equal(t, "bar", attributes["http.request.header.chargebee-foo"])
}

func TestBuildRequestHeaderSpanAttributesExcludesRequestOriginHeaders(t *testing.T) {
	attributes := BuildRequestHeaderSpanAttributes(map[string]string{
		"chargebee-foo":               "bar",
		"chargebee-request-origin-ip": "202.170.207.70",
	})

	require.Equal(t, "bar", attributes["http.request.header.chargebee-foo"])
	require.NotContains(t, attributes, "http.request.header.chargebee-request-origin-ip")
}

func TestBuildRequestHeaderSpanAttributesNormalizesHeaderNames(t *testing.T) {
	attributes := BuildRequestHeaderSpanAttributes(map[string]string{
		"Chargebee-Foo": "bar",
	})

	require.Equal(t, "bar", attributes["http.request.header.chargebee-foo"])
}

func TestBuildRequestHeaderSpanAttributesIgnoresNonChargebeeHeaders(t *testing.T) {
	attributes := BuildRequestHeaderSpanAttributes(map[string]string{
		"Authorization": "Basic secret",
		"X-Custom":      "nope",
	})

	require.Empty(t, attributes)
}
