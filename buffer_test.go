package noticeio

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestDefaultBuffer(t *testing.T) {
	buf := NewBuffer()

	value := []byte("test")
	_, err := buf.Write(value)
	if err != nil {
		t.Fatalf("want no error, got %v", err)
	}
	if err := <-buf.WriteCh; err != nil {
		t.Errorf("want no error, got %v", err)
	}

	dst := make([]byte, len(value))
	_, err = buf.Read(dst)
	if err != nil {
		t.Fatalf("want no error, got %v", err)
	}
	if err := <-buf.ReadCh; err != nil {
		t.Errorf("want no error, got %v", err)
	}
	if !reflect.DeepEqual(value, dst) {
		t.Errorf("want %s, got %s", value, dst)
	}
}

func TestWriteNotice(t *testing.T) {
	buf := NewBufferWithChannel(nil, make(chan error, 1))

	value := []byte("test")
	_, err := buf.Write(value)
	if err != nil {
		t.Fatalf("want no error, got %v", err)
	}
	<-buf.WriteCh

	for i := 0; i < 2; i++ {
		// ignore receive channel
		_, err := buf.Read(nil)
		if err != nil {
			t.Fatalf("want no error, got %v", err)
		}
	}
}

func TestReadNotice(t *testing.T) {
	buf := NewBufferWithChannel(make(chan error, 1), nil)

	for i := 0; i < 2; i++ {
		// ignore receive channel
		value := []byte("test")
		_, err := buf.Write(value)
		if err != nil {
			t.Fatalf("want no error, got %v", err)
		}
	}

	// ignore receive channel
	_, err := buf.Read(nil)
	if err != nil {
		t.Fatalf("want no error, got %v", err)
	}
	<-buf.ReadCh
}

func ExampleBuffer() {
	// notify only write
	buf := NewBufferWithChannel(nil, make(chan error, 1))

	go func(w io.Writer) {
		for {
			w.Write([]byte("example"))
		}
	}(buf)

	// block channel
	err := <-buf.WriteCh
	fmt.Println(err)
	// Output:
	// <nil>
}
