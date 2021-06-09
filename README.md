# go-endpoint

[![Build Status](https://github.com/beyondstorage/go-endpoint/workflows/Unit%20Test/badge.svg?branch=master)](https://github.com/beyondstorage/go-endpoint/actions?query=workflow%3A%22Unit+Test%22)
[![Go dev](https://pkg.go.dev/badge/github.com/beyondstorage/go-endpoint/v4)](https://pkg.go.dev/github.com/beyondstorage/go-endpoint/v4)
[![License](https://img.shields.io/badge/license-apache%20v2-blue.svg)](https://github.com/beyondstorage/go-endpoint/blob/master/LICENSE)
[![go storage dev](https://img.shields.io/matrix/go-endpoint:aos.dev.svg?server_fqdn=chat.aos.dev&label=%23go-endpoint%3Aaos.dev&logo=matrix)](https://matrix.to/#/#go-endpoint:aos.dev) <!-- Need update after matrix updated -->

Both human and machine readable  endpoint format.

## Format

```
<protocol>:<value>+
```

For example:

- File: `file:/var/cache/data`
- HTTP: `http:example.com:80`
- HTTPS: `https:example.com:443`

## Quick Start

```go
ep, err := endpoint.Parse("https:example.com")
if err != nil {
	log.Fatal("parse: ", err)
}

ep.Protocol() // -> "https"
ep.HTTPS() // -> "https://example.com", "example.com", 443
ep.HTTP() // -> panic
ep.File() // -> panic
```
