package observer

import "sync"

// Property is an object that is continuously updated by one or more
// publishers.  It is completely goroutine safe: you can use Property concurrently
// from multiple goroutines.
type Property interface {
	// Value returns the current value for this property.
	Value() interface{}

	// Update sets a new value for this property.
	Update(value interface{})

	// Observe returns a newly created Stream for this property.
	Observe() Stream
}

// NewProperty creates a new Property with initial value value.
// It returns the created Property.
func NewProperty(value interface{}) Property {
	return &property{state: newState(value)}
}

type property struct {
	sync.RWMutex
	state *state
}

func (p *property) Value() interface{} {
	p.RLock()
	p.RUnlock()
	return p.state.value
}

func (p *property) Update(value interface{}) {
	p.Lock()
	p.Unlock()
	p.state = p.state.update(value)
}

func (p *property) Observe() Stream {
	p.RLock()
	p.RUnlock()
	return &stream{state: p.state}
}