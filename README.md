# lib-utils-golang

Small Go helpers we kept rewriting across Keytiles libs — placeholders, maps, pointers, and readable dumps.

Package: [`pkg/kt_utils`](pkg/kt_utils) · module `github.com/keytiles/lib-utils-golang/v2`

## What's in the box

More detail is in versioned files named `<Feature>-vX.Y.md` inside [`docs/`](docs/) (open the highest `vX.Y` for each topic).

- **Strings** — resolve / extract `{var}` placeholders (`StringFunctions-…`)
- **Maps** — keys as a slice, values as a set (`MapsFunctions-…`)
- **Pointers** — nil-safe deref (`ValueFromPtr`); `Ptr` deprecated on Go 1.26+ (`SimpleFuncs-…`)
- **Printing** — litter-backed dumps + lazy `VarPrinter` for logs (`VarPrinting-…`)

## How to use?

```go
import "github.com/keytiles/lib-utils-golang/v2/pkg/kt_utils"
```

## Changes

See [CHANGELOG.md](CHANGELOG.md) for release history.
