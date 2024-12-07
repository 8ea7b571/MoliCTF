package mCache

import "sync"

type MCache struct {
	User *userCache
}

/*
	User Cache
*/

type userCache struct {
	cache *sync.Map
}

func (uc *userCache) Set(key string, value interface{}) {
	uc.cache.Store(key, value)
}

func (uc *userCache) Get(key string) interface{} {
	if value, ok := uc.cache.Load(key); ok {
		return value
	} else {
		return nil
	}
}

func (uc *userCache) Del(key string) {
	uc.cache.Delete(key)
}

func NewMCache() *MCache {
	uCache := &userCache{
		cache: new(sync.Map),
	}

	return &MCache{
		User: uCache,
	}
}
