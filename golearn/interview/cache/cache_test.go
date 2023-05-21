/*
*

	@author: jiangxi
	@since: 2023/5/21
	@desc: //TODO

*
*/
package cache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := NewMemCache()
	cache.SetMaxMemory("100MB")
	cache.Set("int", 1, time.Second)
	cache.Set("bool", false)
	cache.Set("data", map[string]interface{}{"a": 1}, time.Second)

	//cache.Set("int", 1)
	//cache.Set("bool", false)
	//cache.Set("data", map[string]interface{}{"a": 1})
	cache.Get("int")
	time.Sleep(time.Second)
	cache.Get("int")
	cache.Get("bool")
	cache.Del("int")
	cache.Get("bool")
	cache.Keys()
	cache.Flush()
	cache.Keys()
}
