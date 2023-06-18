/*
*

	@author: jiangxi
	@since: 2022/11/11
	@desc: //TODO

*
*/
package Trie

//设计一个 map ，满足以下几点:
//
//字符串表示键，整数表示值
//返回具有前缀等于给定字符串的键的值的总和
//实现一个 MapSum 类：
//
//MapSum() 初始化 MapSum 对象
//void insert(String key, int val) 插入 key-val 键值对，字符串表示键 key ，
//整数表示值 val 。如果键 key 已经存在，那么原来的键值对 key-value 将被替代成新的键值对。
//int sum(string prefix) 返回所有以该前缀 prefix 开头的键 key 的值的总和。

type TrieNode struct {
	children [26]*TrieNode
	val      int //每个节点都保存了值，就是一个根值最大的结构
}

type MapSum struct {
	root *TrieNode
	cnt  map[string]int
}

func Constructor() MapSum {
	return MapSum{&TrieNode{}, map[string]int{}}
}

func (m *MapSum) Insert(key string, val int) {
	delta := val
	if m.cnt[key] > 0 {
		delta -= m.cnt[key] //有可能已经插入过
	}
	m.cnt[key] = val
	node := m.root
	for _, ch := range key {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &TrieNode{}
		}
		node = node.children[ch]
		node.val += delta
	}
}

func (m *MapSum) Sum(prefix string) int {
	node := m.root
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return 0
		}
		node = node.children[ch]
	}
	return node.val
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
