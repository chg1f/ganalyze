package ganalyze

import "time"

type Context[T any] struct {
	StartAt time.Time
	StopAt  time.Time
	Value   T
}

func (c *Context[T]) Used() time.Duration {
	stopAt := time.Now()
	if !c.StopAt.IsZero() {
		stopAt = c.StopAt
	}
	return stopAt.Sub(c.StartAt)
}

type HandleContext[T any] func(c *Context[T])

var (
	EnableAnalyze bool
)

func Start[T any](handles ...HandleContext[T]) *Context[T] {
	if EnableAnalyze {
		startAt := time.Now()
		c := new(Context[T])
		c.StartAt = startAt
		for _, handle := range handles {
			handle(c)
		}
		return c
	}
	return nil
}
func (c *Context[T]) Stop(handles ...HandleContext[T]) {
	if c != nil && EnableAnalyze {
		stopAt := time.Now()
		c.StopAt = stopAt
		for _, handle := range handles {
			handle(c)
		}
	}
}

func WithStartAt[T any](startAt time.Time) HandleContext[T] {
	return func(c *Context[T]) {
		c.StartAt = startAt
	}
}
func WithStopAt[T any](stopAt time.Time) HandleContext[T] {
	return func(c *Context[T]) {
		c.StopAt = stopAt
	}
}
