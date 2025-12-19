//go:build windows
// +build windows

package signals

import (
	"context"
	"os"
)

// SetupSIGHUPHandler returns a dummy channel for Windows as SIGHUP is not supported
func SetupSIGHUPHandler(ctx context.Context) <-chan os.Signal {
	return make(chan os.Signal)
}
