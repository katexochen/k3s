//go:build !windows
// +build !windows

package signals

import (
	"context"
	"os"
	"os/signal"

	"golang.org/x/sys/unix"
)

// SetupSIGHUPHandler registers for SIGHUP. A channel is returned
// which receives the signal.
func SetupSIGHUPHandler(ctx context.Context) <-chan os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, unix.SIGHUP)

	go func() {
		<-ctx.Done()
		signal.Stop(ch)
		close(ch)
	}()

	return ch
}
