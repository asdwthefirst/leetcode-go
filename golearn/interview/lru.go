package interview

import (
	"container/list"
	"fmt"
)

// 在这个实现中，我们使用了一个双向链表（list.List）来维护键值对的顺序，
// 以及一个字典（map）来快速查找键值对。当我们访问一个键值对时，
// 我们将其移动到链表的前面，以表示它是最近访问的。当我们需要删除一个键值对时，
// 我们删除链表的最后一个元素，并从字典中删除相应的键值对。
type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type entry struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.cache[key]; ok {
		this.list.MoveToFront(elem)
		return elem.Value.(*entry).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if elem, ok := this.cache[key]; ok {
		this.list.MoveToFront(elem)
		elem.Value.(*entry).value = value
		return
	}
	if this.list.Len() == this.capacity {
		tail := this.list.Back()
		delete(this.cache, tail.Value.(*entry).key)
		this.list.Remove(tail)
	}
	elem := this.list.PushFront(&entry{key, value})
	this.cache[key] = elem
}

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // Output: 1
	cache.Put(3, 3)
	fmt.Println(cache.Get(2)) // Output: -1
	cache.Put(4, 4)
	fmt.Println(cache.Get(1)) // Output: -1
	fmt.Println(cache.Get(3)) // Output: 3
	fmt.Println(cache.Get(4)) // Output: 4
}
