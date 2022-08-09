//go:build tcp

package main

import (
	"io"
	"net"
	"time"
	"unsafe"
)

func initBridge(b interface {
	Error(...any)
	Fatal(...any)
}, fn ...string) Bridge {
	con, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP: net.IPv4(127, 0, 0, 1), Port: 13000,
	})
	if err != nil {
		b.Fatal(err)
	}

	con.SetWriteBuffer(9_000_000)
	con.SetReadBuffer(9_000_000)
	con.SetDeadline(time.Now().Add(time.Hour))

	return &TCP{
		conn: con,
	}
}

type TCP struct {
	conn *net.TCPConn
	out  []byte
	buf  []byte
}

func (t *TCP) Write(b []byte) bool {
	t.out = b
	return true
}

func (t *TCP) Reader(l uint32) []byte {
	if cap(t.buf) < int(l) {
		t.buf = make([]byte, l)
	}
	if _, err := io.ReadAtLeast(t.conn, t.buf, int(l)); err != nil {
		panic(err)
	}
	return t.buf
}

func (t *TCP) ReaderReset(b []byte) {
	for i := range t.buf {
		t.buf[i] = 0
	}
}

func (t *TCP) Run(s Functions, v ...uint64) ([]uint64, error) {
	lout := len(t.out)
	t.conn.Write((*[4]byte)(unsafe.Pointer(&lout))[:])
	t.conn.Write((*[4]byte)(unsafe.Pointer(&s))[:])
	t.conn.Write(t.out)
	t.out = nil

	out := make(chan []uint64)
	go func() {
		switch s {
		case FunctionKBenchmarkDecodeObjectAPI:
			o := make([]byte, 4)
			if _, err := io.ReadAtLeast(t.conn, o, 4); err != nil {
				panic(err)
			}
			if _, err := io.ReadAtLeast(t.conn, t.out, int(*(*uint32)(unsafe.Pointer(&o[0])))); err != nil {
				panic(err)
			}
			out <- nil
		case FunctionKBenchmarkDecodeSumVec3, FunctionKBenchmarkEncodeObjectAPI:
			o := make([]byte, 4)
			if _, err := io.ReadAtLeast(t.conn, o, 4); err != nil {
				panic(err)
			}
			out <- []uint64{uint64(*(*uint32)(unsafe.Pointer(&o[0])))}
		}
	}()

	return <-out, nil
}

func (t *TCP) Close() error {
	return t.conn.Close()
}
