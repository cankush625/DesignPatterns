package object_pool_design_pattern

import (
	"fmt"
	"sync"
)

type pool struct {
	idle   []IPoolObject
	active []IPoolObject
	capacity int
	mulock   *sync.Mutex
}

// initPool initializes the pool
func initPool(poolObjects []IPoolObject) (*pool, error) {
	if len(poolObjects) == 0 {
		return nil, fmt.Errorf("cannot craete a pool of length 0")
	}
	active := make([]IPoolObject, 0)
	pool := &pool{
		idle:     poolObjects,
		active:   active,
		capacity: len(poolObjects),
		mulock:   new(sync.Mutex),
	}
	return pool, nil
}

// loan gets the pool object from object pool idle objects
func (p *pool) loan() (IPoolObject, error) {
	p.mulock.Lock()
	defer p.mulock.Unlock()
	if len(p.idle) == 0 {
		return nil, fmt.Errorf("all pool objects are occupied. Please request after sometime")
	}
	obj := p.idle[0]
	p.idle = p.idle[1:]
	p.active = append(p.active, obj)
	fmt.Printf("Loan Pool Object with ID: %s\n", obj.getID())
	return obj, nil
}

// release returns the pool object in use back to the Object Pool
func (p *pool) release(target IPoolObject) error {
	p.mulock.Lock()
	defer p.mulock.Unlock()
	err := p.remove(target)
	if err != nil {
		return err
	}
	p.idle = append(p.idle, target)
	fmt.Printf("Return Pool Object with ID: %s\n", target.getID())
	return nil
}

// remove the pool object from the active poll objects list
func (p *pool) remove(target IPoolObject) error {
	currentActiveLength := len(p.active)
	for i, obj := range p.active {
		if obj.getID() == target.getID() {
			p.active[currentActiveLength-1], p.active[i] = p.active[i], p.active[currentActiveLength-1]
			p.active = p.active[:currentActiveLength-1]
			return nil
		}
	}
	return fmt.Errorf("targe pool object doesn't belong to the pool")
}
