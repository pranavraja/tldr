package tldr

import (
	"bytes"
	"io"
)

func NewCachedPage(contents []byte) Page {
	return &CachedPage{
		contents: contents,
	}
}

type CachedPage struct {
	contents []byte
}

func (p *CachedPage) Reader() io.Reader {
	return bytes.NewBuffer(p.contents)
}

func (p *CachedPage) Close() error {
	return nil
}
