package cache

import (
	"errors"
	"log"
	"strconv"
	"sync"
)

type Fnc int

const (
	Get Fnc = iota
	Set
	Delete
)

type Params struct {
	FncType Fnc    `json:"fnc"`
	Key     string `json:"key"`
	Val     any    `json:"val"`
}

type Response struct {
	Hit  bool
	Data any
}

// Generic in-mem cache for now.

type Cache interface {
	Call(params, data any) error
	Get(key string) (any, bool)
	Set(key string, value any)
	Delete(key string)
}

type cache struct {
	mu    sync.RWMutex
	items map[string]any
}

func New() *cache {
	return &cache{
		items: map[string]any{},
		mu:    sync.RWMutex{},
	}
}

// The way I'm going about this seems a bit convoluted. Might change this pattern at some stage.
func (c *cache) Call(params any) (any, error) {
	log.Println("calling cache service")

	pMap, ok := params.(map[string]any)
	if !ok {
		return nil, errors.New("invalid cache params")
	}

	// Would validate these or convert some other way.
	var p Params

	fncVal, ok := pMap["fnc"].(string)
	if !ok {
		p.FncType = -1
	}

	fncInt, err := strconv.Atoi(fncVal)
	if err != nil {
		return nil, err
	}

	p.FncType = Fnc(fncInt)

	key, ok := pMap["key"].(string)
	if ok {
		p.Key = key
	}
	val, found := pMap["val"]
	if found {
		p.Val = val
	}

	switch p.FncType {
	case Get:
		val, found := c.Get(p.Key)
		return &Response{Data: val, Hit: found}, nil
	case Set:
		c.Set(p.Key, p.Val)
		return nil, nil
	case Delete:
		c.Delete(p.Key)
		return nil, nil
	default:
	}

	return nil, errors.New("invalid cache fnc")
}

func (c *cache) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[key]
	return item, found
}

func (c *cache) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = value
}

func (c *cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
}
