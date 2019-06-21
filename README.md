# three

three.js wrapper for GopherJS or WASM

This 'Go' code wrapper generated from Typedoc JSON.

# Re Generate

```shell
cd internal/generate
git -C three.js checkout r104
make
```

# USAGE

|JavaScript|GopherJS|
|--|--|
|new THREE.Scene()|core.NewScene()|

# SAMPLE

```go
package main

import (
	"github.com/nobonobo/three/core"
)

func main() {
	scene := core.NewScene()
}
```

