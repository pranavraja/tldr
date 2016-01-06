package remote

import "io"

func NewRemotePage(readCloser io.ReadCloser) *RemotePage {
	return &RemotePage{
		readCloser: readCloser,
	}
}

type RemotePage struct {
	readCloser io.ReadCloser
}

func (p *RemotePage) Reader() io.Reader {
	return p.readCloser
}

func (p *RemotePage) Close() error {
	return p.readCloser.Close()
}
