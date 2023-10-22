# `system-shutdown`

[![CI/CD][ci-cd-badge]][ci-cd-url]

`system-shutdown` provides a cross-platform way to shut down, reboot, sleep, or hibernate operations.

Supported platforms: Linux, Windows, and macOS.

Based on the work from the Rust crate [`system_shutdown`](https://github.com/risoflora/system_shutdown).

## Usage

Add this to your `go.mod`:

```
require (
    github.com/plackemacher/system-shutdown
)
```

and then:

```go
package main

import (
    "fmt"
    "github.com/plackemacher/system-shutdown"
    "os"
)

func main() {
    err := system_shutdown.Shutdown()
    if err != nil {
        println("Shutting down, bye!")
    } else {
        fmt.Fprintf(os.Stderr, "Failed to shut down: %v", err)
    }
}
```

In most of the systems it does not require the user to be root or admin.

## Contributions

Pull Requests are welcome! =)

## License

`system-shutdown` is licensed under either of the following, at your option:

- [Apache License 2.0](LICENSE-APACHE)
- [MIT License](LICENSE-MIT)

[ci-cd-badge]: https://github.com/plackemacher/system-shutdown/actions/workflows/CI.yml/badge.svg
[ci-cd-url]: https://github.com/plackemacher/system-shutdown/actions/workflows/CI.yml
[license-url]: https://github.com/plackemacher/system-shutdown#license