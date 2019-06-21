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

| JavaScript        | GopherJS         |
| ----------------- | ---------------- |
| new THREE.Scene() | three.NewScene() |

# SAMPLE

```go
package main

import (
	"fmt"

	"github.com/nobonobo/three"
)

func main() {
	scene := three.NewScene()
	fmt.Println(scene)
}
```
