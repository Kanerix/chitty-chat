package mvc

import "errors"

const (
	Width = 30
)

var (
	ErrInvalidMessage = errors.New("invalid message")
	ErrStreamClosed   = errors.New("stream closed")
)
