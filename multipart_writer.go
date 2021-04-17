package coreutils

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/aos-dev/go-storage/v3/types"
)

// MultipartWriter is used to write multiparts.
type MultipartWriter struct {
	m     types.Multiparter
	o     *types.Object
	parts []*types.Part

	idx int
}

// NewMultipartWriter will create a new multipart writer.
// If input storager doesn't implement Multiparter, ErrMultiparterNotImplemented will be returned.
func NewMultipartWriter(ctx context.Context, s types.Storager, path string, ps ...types.Pair) (mw *MultipartWriter, err error) {
	m, ok := s.(types.Multiparter)
	if !ok {
		return nil, ErrMultiparterNotImplemented
	}

	o, err := m.CreateMultipartWithContext(ctx, path, ps...)
	if err != nil {
		return nil, fmt.Errorf("create_multipart: %w", err)
	}

	mw = &MultipartWriter{
		m:     m,
		o:     o,
		parts: nil,
		idx:   0,
	}

	return mw, nil
}

// Write will write bytes as a part.
//
// NOTES:
//   - Write is not concurrent safe.
func (m MultipartWriter) Write(ctx context.Context, p []byte) (n int, err error) {
	length := int64(len(p))

	nn, err := m.m.WriteMultipartWithContext(ctx, m.o, bytes.NewReader(p), length, m.idx)
	if err != nil {
		return int(nn), fmt.Errorf("write_multiaprt: %w", err)
	}

	m.parts = append(m.parts, &types.Part{
		Index: m.idx,
		Size:  nn,
		ETag:  "",
	})
	m.idx++
	return int(nn), nil
}

// ReadFrom will write bytes from an io.Reader as a part.
//
// NOTES:
//   - Write is not concurrent safe.
func (m MultipartWriter) ReadFrom(ctx context.Context, r io.Reader, l int64) (n int64, err error) {
	n, err = m.m.WriteMultipartWithContext(ctx, m.o, r, l, m.idx)
	if err != nil {
		return n, fmt.Errorf("write_multiaprt: %w", err)
	}

	m.parts = append(m.parts, &types.Part{
		Index: m.idx,
		Size:  n,
		ETag:  "",
	})
	m.idx++
	return n, nil
}

// Close will complete multipart.
func (m MultipartWriter) Close(ctx context.Context) error {
	err := m.m.CompleteMultipartWithContext(ctx, m.o, m.parts)
	if err != nil {
		return fmt.Errorf("complete_multipart: %w", err)
	}
	return nil
}
