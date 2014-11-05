// Copyright 2014 li. All rights reserved.
// Use of this source code is governed by a MIT/X11
// license that can be found in the LICENSE file.

package containers

import (
	"sync"
)

// Get a new ConcurrentMap instance.
func NewCMap() ConcurrentMap {
	return &concMap{elements: make(map[interface{}]interface{})}
}

// Threadsafe Map.
type ConcurrentMap interface {

	// Maps the specified key to the specified value.
	// Neither the key nor the value can be nil.
	// The value can be retrieved by calling the get method with a key that is equal to the original key.
	Put(k interface{}, v interface{})

	// If the specified key is not already associated with a value,
	// associate it with the given value. This is equivalent to
	// if (!map.containsKey(key))
	//   return map.put(key, value);
	// else
	//   return map.get(key);
	PutIfAbsent(k interface{}, v interface{})

	// Returns the value to which the specified key is mapped,
	// isExists value indicates whether this map contains mapping for the key.
	Get(k interface{}) (v interface{}, isExists bool)

	// Check if the map contains the key.
	ContainsKey(k interface{}) (isExists bool)

	// Return an slice of the keys in this container.
	Keys() []interface{}

	// Removes all of the mappings from this map.
	Clear()

	// Removes the key (and its corresponding value) from this map.
	Remove(k interface{})

	// Returns the number of key-value mappings in this map.
	Size() int
}

type concMap struct {
	elements map[interface{}]interface{}
	mutex    sync.RWMutex
}

func (c *concMap) Put(k interface{}, v interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.elements[k] = v
}

func (c *concMap) PutIfAbsent(k interface{}, v interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	_, isExists := c.elements[k]
	if !isExists {
		c.elements[k] = v
	}
}

func (c *concMap) Get(k interface{}) (v interface{}, isExists bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	v, isExists = c.elements[k]
	return
}

func (c *concMap) ContainsKey(k interface{}) (isExists bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	_, isExists = c.elements[k]
	return isExists
}

func (c *concMap) Keys() []interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	keys := make([]interface{}, len(c.elements))
	i := 0
	for k, _ := range c.elements {
		keys[i] = k
		i++
	}
	return keys
}

func (c *concMap) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.elements = make(map[interface{}]interface{})
}

func (c *concMap) Remove(k interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.elements, k)
}

func (c *concMap) Size() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return len(c.elements)
}
