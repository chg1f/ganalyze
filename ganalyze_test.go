package ganalyze

import (
	"fmt"
	"strconv"
	"sync/atomic"
	"testing"
	"time"
)

func TestCount(t *testing.T) {
	EnableAnalyze = true
	defer Start(func(c *Context[uint64]) {
		fmt.Println("Count: " + strconv.FormatUint(atomic.LoadUint64(&c.Value), 10))
		fmt.Println("Used: " + c.Used().String())
		if atomic.LoadUint64(&c.Value) != 0 {
			t.Error(c.Value)
		}
		atomic.AddUint64(&c.Value, 1)
	}).Stop(func(c *Context[uint64]) {
		fmt.Println("Count: " + strconv.FormatUint(atomic.LoadUint64(&c.Value), 10))
		fmt.Println("Used: " + c.Used().String())
		if c.Used().Truncate(time.Second) != time.Second*2 {
			t.Error(c.Used().String())
		}
		if atomic.LoadUint64(&c.Value) != 1 {
			t.Error(c.Value)
		}
	})
	time.Sleep(time.Second * 2)
}

func Example_Count(t *testing.T) {
	EnableAnalyze = true
	defer Start(func(c *Context[uint64]) {
		fmt.Println("Start Count: " + strconv.FormatUint(atomic.LoadUint64(&c.Value), 10)) // 0
	}).Stop(func(c *Context[uint64]) {
		fmt.Println("Stop Count: " + strconv.FormatUint(atomic.LoadUint64(&c.Value), 10)) // 1
		fmt.Println("Stop Used: " + c.Used().Truncate(time.Second).String())              // 2s
	})
	time.Sleep(time.Second * 2)
	// Output:
	// Start Count: 0
	// Stop Count: 1
	// Stop Count: 2s
}
