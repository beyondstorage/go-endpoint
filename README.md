# go-endpoint

Both human and machine readable  endpoint format.

## Notes

**This package has been moved to [go-storage](https://github.com/beyondstorage/go-storage/tree/master/endpoint).**

```shell
go get go.beyondstorage.io/endpoint
```

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

switch ep.Protocol() {
case ProtocolHTTP:
    url, host, port := ep.HTTP()
    log.Println("url: ", url)
    log.Println("host: ", host)
    log.Println("port: ", port)
case ProtocolHTTPS:
    url, host, port := ep.HTTPS()
    log.Println("url: ", url)
    log.Println("host: ", host)
    log.Println("port: ", port)
case ProtocolFile:
    path := ep.File()
    log.Println("path: ", path)
default:
    panic("unsupported protocol")
}
```
