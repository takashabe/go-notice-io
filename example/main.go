package main

import "io"

func main() {
	// notify only write
	buf := noticeio.NewBufferWithChannel(nil, make(chan error, 1))

	go func(w io.Writer) {
		for {
			w.Write([]byte("example"))
		}
	}(buf)

	// block channel
	<-buf.WriteCh
}
