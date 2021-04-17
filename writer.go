package coreutils

import (
	"context"
	"io"
)

// Writer is a high level abstraction for go-storage's Multiparter, Appender, ...
type Writer interface {
	Write(ctx context.Context, p []byte) (n int, err error)
	ReadFrom(ctx context.Context, r io.Reader, l int64) (n int64, err error)
	Close(ctx context.Context) error
}

func init() {
	var _ Writer = AppendWriter{}
	var _ Writer = MultipartWriter{}
}
