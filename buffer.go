// Package noticeio implements a wrapper for the io package and add notification via channel.
package noticeio

import "bytes"

// Buffer is wrapped bytes.Buffer and notify channel at call Read and Write method
type Buffer struct {
	buf *bytes.Buffer

	// read channel
	internalRead chan error
	ReadCh       <-chan error

	// write channel
	internalWrite chan error
	WriteCh       <-chan error
}

// NewBuffer return new NoticeWriter with default channel
func NewBuffer() *Buffer {
	r := make(chan error, 1)
	w := make(chan error, 1)
	return NewBufferWithChannel(r, w)
}

// NewBuffer return new NoticeWriter with specific channel
func NewBufferWithChannel(r chan error, w chan error) *Buffer {
	return &Buffer{
		buf:           new(bytes.Buffer),
		internalRead:  r,
		ReadCh:        r,
		internalWrite: w,
		WriteCh:       w,
	}
}

// Write call internal write method and send error to channel
func (b *Buffer) Write(p []byte) (n int, err error) {
	n, err = b.buf.Write(p)
	if b.internalWrite != nil {
		b.internalWrite <- err
	}
	return
}

// Read call internal read method and send error to channel
func (b *Buffer) Read(p []byte) (n int, err error) {
	n, err = b.buf.Read(p)
	if b.internalRead != nil {
		b.internalRead <- err
	}
	return
}
