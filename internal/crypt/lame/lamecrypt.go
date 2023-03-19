package lame

import (
	"io"
)

var cryptString string = "Range check error while converting variant of type (%s) into type (%s)"

func wrapInput(input io.Reader) io.Reader {
	return &inputWrapper{input: input}
}

type inputWrapper struct {
	input io.Reader
	pos   int
}

func (w *inputWrapper) Read(p []byte) (n int, err error) {
	n, err = w.input.Read(p)

	for i := 0; i < n; i++ {
		p[i] ^= cryptString[w.pos]
		w.pos++
		if w.pos == len(cryptString) {
			w.pos = 0
		}
	}

	return n, err
}

func wrapOutput(output io.Writer) io.Writer {
	return &outputWrapper{output: output}
}

type outputWrapper struct {
	output io.Writer
	pos    int
}

func (w *outputWrapper) Write(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		p[i] ^= cryptString[w.pos]
		w.pos++
		if w.pos == len(cryptString) {
			w.pos = 0
		}
	}

	return w.output.Write(p)
}
