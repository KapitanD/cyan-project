package utils

import (
	"context"
	"errors"
	"sync"
	"time"
)

type TTLMapI interface {
	Set(key string, value interface{}, ttl time.Duration)
	Get(key string) (interface{}, error)
	Del(Key string) error
}

type item struct {
	value interface{}
	exp   time.Time
	mu    *sync.RWMutex
}

func NewTTLMap(ctx context.Context) *TTLMap {
	ttlmap := &TTLMap{
		m: make(map[string]*item),
	}

	go ttlmap.cleanLoop(ctx)
	return ttlmap
}

type TTLMap struct {
	m map[string]*item
}

func (ttlm *TTLMap) clean() {
	for key, item := range ttlm.m {
		ttlm.m[key].mu.Lock()

		if time.Now().After(item.exp) {
			delete(ttlm.m, key)
		}

		ttlm.m[key].mu.Unlock()
	}
}

func (ttlm *TTLMap) cleanLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			ttlm.clean()
		}
		time.Sleep(time.Second)
	}
}

func (ttlm *TTLMap) Set(key string, value interface{}, ttl time.Duration) {
	if _, ok := ttlm.m[key]; !ok {
		ttlm.m[key] = &item{
			value: value,
			exp:   time.Now().Add(ttl),
			mu:    &sync.RWMutex{},
		}
		return
	}

	ttlm.m[key].mu.Lock()
	defer ttlm.m[key].mu.Unlock()
	ttlm.m[key].value = value
	ttlm.m[key].exp = time.Now().Add(ttl)
}

func (ttlm *TTLMap) Get(key string) (interface{}, error) {
	ttlm.m[key].mu.RLock()
	defer ttlm.m[key].mu.RUnlock()

	if _, ok := ttlm.m[key]; !ok {
		return nil, errors.New("no such key")
	}

	val := ttlm.m[key].value
	return val, nil
}

func (ttlm *TTLMap) Del(key string) error {
	ttlm.m[key].mu.Lock()
	defer ttlm.m[key].mu.Unlock()

	if _, ok := ttlm.m[key]; !ok {
		return errors.New("no such key")
	}

	delete(ttlm.m, key)

	return nil
}
