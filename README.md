# go-notice-io

go-notice-io is `io` package wrapper.
Notify via channel each time when call Write and Read method.

This is useful if you want to test writing to io.Writer in the goroutine and so on.

# Installation

```
$ go get github.com/takashabe/go-notice-io
```

# Usage

```go
package main

import (
  "io"

  "github.com/takashabe/go-notice-io"
)

func main() {
  // notify only write
  buf := noticeio.NewBufferWithChannel(nil, make(chan error, 1))

  go func(w io.Writer) {
    for {
      w.Write([]byte("example"))
    }
  }(buf)

  // receive channel each time Write()
  <-buf.WriteCh
}
```
