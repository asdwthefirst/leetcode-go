/*
*

	@author: jiangxi
	@since: 2023/5/21
	@desc: //TODO

*
*/
package cache

import (
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type MemCache struct {
	//最大内存，单位Byte
	maxMemory int64
	//已使用内存
	currentUsedMemory int64
	//键值对
	values map[string]memCacheValue
	//键值对锁
	valuesMutex sync.RWMutex
	//开关
	shutdownCh chan struct{}
}

type memCacheValue struct {
	val      interface{}
	size     int64
	expireAt time.Time
	expire   bool
}

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
)

// 设置最大大小，整数+KB,MB,GB,TB

// TODO:setmaxmemory也要考虑安全问题。
func (cache *MemCache) SetMaxMemory(maxSize string) bool {
	//正则表达式
	re, _ := regexp.Compile("[0-9]+")
	sizeNum, _ := strconv.ParseInt(string(re.Find([]byte(maxSize))), 10, 64)
	sizeUnit := strings.ToUpper(string(re.ReplaceAll([]byte(maxSize), []byte(""))))
	var size int64
	switch sizeUnit {
	case "KB":
		size = sizeNum * KB
	case "MB":
		size = sizeNum * MB
	case "GB":
		size = sizeNum * GB
	default:
		size = 0
	}
	if size <= 0 {
		return false
	}
	cache.maxMemory = size
	return true
}

// 设置值
func (cache *MemCache) Set(key string, val interface{}, expire ...time.Duration) bool {
	cache.valuesMutex.Lock()
	defer cache.valuesMutex.Unlock()
	if _, ok := cache.get(key); ok {
		cache.del(key)
	}
	cache.add(key, val, expire...)
	if cache.currentUsedMemory > cache.maxMemory {
		cache.del(key)
		return false
	}
	return true
}

func (cache *MemCache) add(key string, val interface{}, expireTime ...time.Duration) bool {
	memVal := memCacheValue{
		val:    val,
		size:   getObjectSize(val),
		expire: len(expireTime) > 0,
	}
	if memVal.expire {
		memVal.expireAt = time.Now().Add(expireTime[0])
	}
	cache.values[key] = memVal
	cache.currentUsedMemory += memVal.size
	return true
}

func (cache *MemCache) del(key string) bool {
	_, ok := cache.values[key]
	if !ok {
		return false
	}
	cache.currentUsedMemory -= cache.values[key].size
	delete(cache.values, key)
	return true
}

// Set和Get涉及读写锁的访问，避免互相调用，添加一个内部不带锁的方法
func (cache *MemCache) get(key string) (interface{}, bool) {
	if val, ok := cache.values[key]; !ok || val.expire && val.expireAt.Before(time.Now()) {
		return nil, false
	} else {
		return val.val, true
	}
}

func (cache *MemCache) Get(key string) (interface{}, bool) {
	cache.valuesMutex.RLock()
	defer cache.valuesMutex.RUnlock()
	return cache.get(key)
}

// 删除值
func (cache *MemCache) Del(key string) bool {
	cache.valuesMutex.Lock()
	defer cache.valuesMutex.Unlock()
	return cache.del(key)
}

// 是否存在值
func (cache *MemCache) Exists(key string) bool {
	cache.valuesMutex.RLock()
	defer cache.valuesMutex.RUnlock()
	_, ok := cache.get(key)
	return ok
}

// 清空内存
func (cache *MemCache) Flush() bool {
	cache.valuesMutex.Lock()
	defer cache.valuesMutex.Unlock()
	for k, _ := range cache.values {
		cache.del(k)
	}
	return false
}

// 获取key的数量
func (cache *MemCache) Keys() int {
	cache.valuesMutex.RLock()
	defer cache.valuesMutex.RUnlock()
	return len(cache.values)
}

func NewMemCache() *MemCache {
	return &MemCache{
		values: make(map[string]memCacheValue),
	}

}

func (cache *MemCache) periodicCleanup(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			cache.cleanup()
		case <-cache.shutdownCh:
			return
		}
	}
}

func (cache *MemCache) cleanup() {
	cache.valuesMutex.Lock()
	defer cache.valuesMutex.Unlock()
	for k, v := range cache.values {
		if v.expire && v.expireAt.After(time.Now()) {
			cache.del(k)
		}
	}
	return
}

func (cache *MemCache) ShutDown() {
	cache.shutdownCh <- struct{}{}
	return
}
