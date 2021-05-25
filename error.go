package coreutils

import (
	"errors"
)

var (
	ErrMultipleWriteNotSupported = errors.New("multiple write not supported")
	ErrMultiparterNotImplemented = errors.New("multiparter not implemented")
	ErrAppenderNotImplemented    = errors.New("appender not implemented")
)
