package coreutils

import (
	"errors"
)

var (
	ErrMultiparterNotImplemented = errors.New("multiparter not implemented")
	ErrAppenderNotImplemented    = errors.New("appender not implemented")
)
