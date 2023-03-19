package xor

import (
	"io"
)

type L2Ver1x1InputStream struct {
	input  io.Reader
	xorKey int
}

func NewL2Ver1x1InputStream(input io.Reader, xorKey int) *L2Ver1x1InputStream {
	return &L2Ver1x1InputStream{
		input:  input,
		xorKey: xorKey,
	}
}

func (is *L2Ver1x1InputStream) Read(p []byte) (n int, err error) {
	n, err = is.input.Read(p)
	if err != nil {
		return 0, err
	}

	for i := 0; i < n; i++ {
		p[i] ^= byte(is.xorKey)
	}

	return n, nil
}

func (s *L2Ver1x1InputStream) ReadByte() (byte, error) {
	b := make([]byte, 1)
	_, err := s.Read(b)
	if err != nil {
		return 0, err
	}

	return b[0], nil
}
