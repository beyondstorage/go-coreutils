package coreutils

import (
	"github.com/beyondstorage/go-storage/v4/services"
)

var (
	ErrMultipleWriteNotSupported = services.NewErrorCode("multiple write not supported")
	ErrMultiparterNotImplemented = services.NewErrorCode("multiparter not implemented")
	ErrAppenderNotImplemented    = services.NewErrorCode("appender not implemented")
)
