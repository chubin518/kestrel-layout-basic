package graceful

import (
	"context"
	"errors"
)

// errHookCallbackExited is returned when a the callback does not finish executing
var errHookCallbackExited = errors.New("goroutine exited without returning")

func withTimeout(ctx context.Context, callback func(context.Context) error) error {
	done := make(chan error, 1)
	go func() {
		// If runtime.Goexit() is called from within the callback
		// then nothing is written to the chan.
		// However the defer will still be called, so we can write to the chan,
		// to avoid hanging until the timeout is reached.
		callbackExited := false
		defer func() {
			if !callbackExited {
				done <- errHookCallbackExited
			}
		}()
		done <- callback(ctx)
		callbackExited = true
	}()

	var err error

	select {
	case <-ctx.Done():
		// returns when callback execution times out
		// context deadline exceeded
		err = ctx.Err()
	case err = <-done:
		// If the context finished at the same time as the callback
		// prefer the context error.
		// This eliminates non-determinism in select-case selection.
		if ctx.Err() != nil {
			err = ctx.Err()
		}
	}
	return err
}
