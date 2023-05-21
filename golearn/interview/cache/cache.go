/*
*

	@author: jiangxi
	@since: 2023/5/21
	@desc: //TODO

*
*/
package cache

import "time"

type Cache interface {
	//设置最大大小，整数+KB,MB,GB,TB
	SetMaxMemory(maxSize string) bool
	//设置值
	Set(key string, val interface{}, expire time.Duration) bool
	//获取值
	Get(key string) (interface{}, bool)
	//删除值
	Del(key string) bool
	//是否存在值
	Exists(key string) bool
	//清空内存
	Flush() bool
	//获取key的数量
	Keys() int64
}
