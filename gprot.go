//this package encodes and decodes in ROT13
package gorot

import "io"

type writer struct {
	wrapped io.Writer
}

type reader struct {
	wrapped io.Reader
}

// NewWriter creates a new io.Writer that encrypts with ROT13.
func NewWriter(wrapped io.Writer) io.Writer {
	return &writer{wrapped}
}

// NewReader creates a new io.Reader that encrypts with ROT13.
func NewReader(wrapped io.Reader) io.Reader {
	return &reader{wrapped}
}

func (r *writer) Write(p []byte) (int, error) {
	o := make([]byte, len(p))
	for i := 0; i < len(p); i++ {
		o[i] = p[i] + 13
	}
	return r.wrapped.Write(o)
}

func (r *reader) Read(p []byte) (int, error) {
	n, err := r.wrapped.Read(p)
	if err != nil {
		return n, err
	}
	//itterate all text and replace
	for i := 0; i < n; i++ {
		p[i] = p[i] - 13
	}
	return n, err
}
