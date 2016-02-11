package cache

import (
	"bytes"
	"io"
)

func NewCachedPage(contents []byte) *CachedPage {
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
