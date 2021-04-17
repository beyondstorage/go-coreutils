package coreutils

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/aos-dev/go-storage/v3/types"
)

// AppendWriter is used to append.
type AppendWriter struct {
	a types.Appender
	o *types.Object
}

func NewAppendWriter(ctx context.Context, s types.Storager, path string, ps ...types.Pair) (aw *AppendWriter, err error) {
	a, ok := s.(types.Appender)
	if !ok {
		return nil, ErrAppenderNotImplemented
	}

	o, err := a.CreateAppendWithContext(ctx, path, ps...)
	if err != nil {
		return nil, fmt.Errorf("create_append: %w", err)
	}

	aw = &AppendWriter{
		a: a,
		o: o,
	}

	return aw, nil
}

// WriteWithContext will write bytes as append.
//
// NOTES:
//   - WriteWithContext is not concurrent safe.
func (a AppendWriter) WriteWithContext(ctx context.Context, p []byte) (n int, err error) {
	nn, err := a.a.WriteAppendWithContext(ctx, a.o, bytes.NewReader(p), int64(len(p)))
	if err != nil {
		return int(nn), fmt.Errorf("write_append: %w", err)
	}
	return int(nn), nil
}

// ReadFromWithContext will write bytes from an io.Reader as append.
//
// NOTES:
//   - ReadFromWithContext is not concurrent safe.
func (a AppendWriter) ReadFromWithContext(ctx context.Context, r io.Reader, l int64) (n int64, err error) {
	n, err = a.a.WriteAppendWithContext(ctx, a.o, r, l)
	if err != nil {
		return n, fmt.Errorf("write_append: %w", err)
	}
	return n, nil
}

// CloseWithContext is a noop for AppendWriter.
func (a AppendWriter) CloseWithContext(ctx context.Context) error {
	return nil
}
