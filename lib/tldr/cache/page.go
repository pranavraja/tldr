package cache

import (
	"bytes"
	"io"

	"github.com/pranavraja/tldr/lib/tldr/entity"
)

func NewCachedPage(contents []byte) entity.Page {
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
