package count

import "sync"

type Counter struct {
	Count int
	mu    sync.Mutex
}

func (c *Counter) Increase() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Count++
}

func (c *Counter) Decrease() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Count--
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.Count
}
