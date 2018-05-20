# Ratchet

Modular transaction engine.

## Usage

As a library:

```go
package main

import (
	"github.com/odinsplasmarifle/ratchet"
)

func main() {
	r := ratchet.NewServer()
	r.Serve()
}
```

As software:

```bash
# Not implemented!
```

## Development

Build the protobuf files:

```bash
protoc -I ratchet/ ratchet/ratchet.proto --go_out=plugins=grpc:ratchet
```
