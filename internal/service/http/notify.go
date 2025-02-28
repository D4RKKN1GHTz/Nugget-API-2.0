//go:build !windows
// +build !windows

package http

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func NotifyContext() (context.Context, context.CancelFunc) {
	return signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, os.Kill)
}
