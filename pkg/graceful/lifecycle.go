package graceful

import "context"

// A Hook is a pair of start and stop callbacks, either of which can be nil,
// plus a string identifying the supplier of the hook.
type Hook struct {
	OnStart func(context.Context) error
	OnStop  func(context.Context) error
}

// Lifecycle coordinates application lifecycle hooks.
type Lifecycle struct {
	hooks []*Hook
}

// Append adds a Hook to the lifecycle.
func (l *Lifecycle) Append(hooks ...*Hook) {
	l.hooks = append(l.hooks, hooks...)
}

// Start runs all OnStart hooks, returning immediately if it encounters an
// error.
func (l *Lifecycle) Start(ctx context.Context) error {
	for _, hook := range l.hooks {
		if err := ctx.Err(); err != nil {
			return err
		}
		if hook.OnStart == nil {
			continue
		}
		if err := hook.OnStart(ctx); err != nil {
			return err
		}
	}
	return nil
}

// Stop runs any OnStop hooks whose OnStart counterpart succeeded. OnStop
// hooks run in reverse order.
func (l *Lifecycle) Stop(ctx context.Context) error {
	numHooks := len(l.hooks)
	for i := numHooks - 1; i >= 0; i-- {
		if err := ctx.Err(); err != nil {
			return err
		}
		hook := l.hooks[i]
		if hook.OnStop == nil {
			continue
		}
		if err := hook.OnStop(ctx); err != nil {
			return err
		}
	}
	return nil
}
