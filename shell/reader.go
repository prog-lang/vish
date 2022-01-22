package shell

import (
	"bufio"
	"io"
	"strings"
)

type Reader struct {
	escaped bool
	reader  *bufio.Reader
	builder strings.Builder
}

func NewReader(r io.Reader) *Reader {
	return &Reader{
		escaped: false,
		reader:  bufio.NewReader(r),
		builder: strings.Builder{},
	}
}

func (r *Reader) Next() (input string, err error) {
	for {
		char, err := r.next()
		if err == io.EOF {
			return r.builder.String(), nil
		}

		if err != nil {
			return "", err
		}

		if char == 0 {
			continue
		}

		if err := r.builder.WriteByte(char); err != nil {
			return "", err
		}
	}
}

func (r *Reader) next() (char byte, err error) {
	if char, err = r.reader.ReadByte(); err != nil {
		return
	}

	if char == '\\' { // '\' is an escape flag
		r.escape()
		return 0, nil
	} else {
		defer r.unescape()
	}

	if char != '\n' {
		return
	}

	if r.escaped { // at this point b == '\n'
		return ' ', nil
	}

	return 0, io.EOF
}

func (r *Reader) escape() {
	r.escaped = true
}

func (r *Reader) unescape() {
	r.escaped = false
}
