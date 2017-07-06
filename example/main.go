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
