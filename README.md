# `fluxx-go`

Official Go SDK for the [Fluxx](https://fluxx.dev) control plane API.

Requires Go 1.22+.

## Install

```bash
go get github.com/fluxx/fluxx-go
```

## Quickstart

```go
package main

import (
    "context"
    "log"
    "github.com/fluxx/fluxx-go"
)

func main() {
    client := fluxx.New("flx_live_…")
    if _, err := client.Health(context.Background()); err != nil {
        log.Fatal(err)
    }
}
```

> **Status:** scaffold. The client surface is being generated from
> `https://api.fluxx.dev/openapi.json`.

## License

Apache-2.0
