package endpoint

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	cases := []struct {
		name  string
		cfg   string
		value Endpoint
		err   error
	}{
		{
			"invalid string",
			"abcx",
			Endpoint{},
			ErrUnsupportedProtocol,
		},
		{
			"normal http",
			"http:example.com:80",
			Endpoint{ProtocolHTTP, hostPort{"example.com", 80}},
			nil,
		},
		{
			"normal http without port",
			"http:example.com",
			Endpoint{ProtocolHTTP, hostPort{"example.com", 80}},
			nil,
		},
		{
			"wrong port number in http",
			"http:example.com:xxx",
			Endpoint{},
			ErrInvalidValue,
		},
		{
			"normal https",
			"https:example.com:443",
			Endpoint{ProtocolHTTPS, hostPort{"example.com", 443}},
			nil,
		},
		{
			"normal https without port",
			"https:example.com",
			Endpoint{ProtocolHTTPS, hostPort{"example.com", 443}},
			nil,
		},
		{
			"wrong port number in https",
			"https:example.com:xxx",
			Endpoint{},
			ErrInvalidValue,
		},
		{
			"not supported protocol",
			"notsupported:abc.com",
			Endpoint{},
			ErrUnsupportedProtocol,
		},
		{
			"normal file",
			"file:/root/data",
			Endpoint{ProtocolFile, "/root/data"},
			nil,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			p, err := Parse(tt.cfg)
			if tt.err == nil {
				assert.Nil(t, err)
			} else {
				assert.True(t, errors.Is(err, tt.err))
			}
			assert.EqualValues(t, tt.value, p)
		})
	}
}

func TestNewFile(t *testing.T) {
	assert.Equal(t, Endpoint{ProtocolFile, "/example"}, NewFile("/example"))
}

func TestNewHTTP(t *testing.T) {
	assert.Equal(t,
		Endpoint{ProtocolHTTP, hostPort{"example.com", 8080}},
		NewHTTP("example.com", 8080),
	)
}

func TestNewHTTPS(t *testing.T) {
	assert.Equal(t,
		Endpoint{ProtocolHTTPS, hostPort{"example.com", 4433}},
		NewHTTPS("example.com", 4433),
	)
}

func TestEndpoint_Protocol(t *testing.T) {
	ep := NewFile("/test")

	assert.Equal(t, ProtocolFile, ep.Protocol())
}

func TestEndpoint_String(t *testing.T) {
	cases := []struct {
		name     string
		value    Endpoint
		expected string
	}{
		{
			"file",
			Endpoint{ProtocolFile, "/test"},
			"file:/test",
		},
		{
			"http without port",
			Endpoint{ProtocolHTTP, hostPort{"example.com", 80}},
			"http:example.com:80",
		},
		{
			"http with port",
			Endpoint{ProtocolHTTP, hostPort{"example.com", 8080}},
			"http:example.com:8080",
		},
		{
			"https without port",
			Endpoint{ProtocolHTTPS, hostPort{"example.com", 443}},
			"https:example.com:443",
		},
		{
			"https with port",
			Endpoint{ProtocolHTTPS, hostPort{"example.com", 4433}},
			"https:example.com:4433",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.value.String())
		})
	}
}

func TestEndpoint(t *testing.T) {
	p := NewFile("/test")

	assert.Panics(t, func() {
		p.HTTP()
	})
	assert.Panics(t, func() {
		p.HTTPS()
	})

	assert.Equal(t, "/test", p.File())
}
