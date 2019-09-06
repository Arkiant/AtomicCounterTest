package benchuuid

import "sync/atomic"

// Counter struct create a new atomic counter
type Counter struct {
	value uint32
}

// New Create a new atomic Counter
func New() Counter {
	return Counter{}
}

// Increment the counter atomically
func (c *Counter) Increment() {
	atomic.AddUint32(&c.value, 1)
}

// Get the current value from the counter
func (c *Counter) Get() uint32 {
	return atomic.LoadUint32(&c.value)
}

// Reset the current value to 0
func (c *Counter) Reset() uint32 {
	return atomic.SwapUint32(&c.value, 0)
}
