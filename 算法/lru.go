package main

import "container/list"

/**
* LRU:最近最少使用，map存放键值对数据，双向链表头部为最近操作过的元素，末尾是最近最少使用的，空间满则移除末尾元素
 */
type node struct{
	key interface{}
	val interface{}
}

type LRU struct{
	cap int
	list *list.List
	cache map[interface{}]*list.Element
}

func NewLRU(cap int) *LRU {
	return &LRU{
		cap: cap,
		list: list.New(),
		cache: make(map[interface{}]*list.Element, cap),
	}
}

func (l *LRU) put(key interface{}, val interface{}) {
	if l.list == nil {
		return
	}
	if ele, ok := l.cache[key]; ok {
		ele.Value.(*node).val = val
		l.list.MoveToFront(ele)
		return
	}

	ele := l.list.PushFront(&node{
		key: key,
		val: val,
	})
	l.cache[key] = ele

	if len(l.cache) > l.cap {
		if e := l.list.Back(); e != nil {
			l.list.Remove(e)
			delete(l.cache, e)
		}

	}

}

func (l *LRU) get(key interface{}) (val interface{}, ok bool){
	if l.cache == nil {
		return
	}

	if ele, ok := l.cache[key]; ok {
		l.list.MoveToFront(ele)
		return ele.Value.(*node).val, true
	}

	return
}

func main() {
	lru := NewLRU(2)
	if v,ok := lru.get(1); ok { println(v.(int))}
	lru.put(1,1)
	lru.put(2,2)
	if v,ok := lru.get(1); ok { println(v.(int))}
	lru.put(3,3)
	if v,ok := lru.get(1); ok { println(v.(int))}
	if v,ok := lru.get(3); ok { println(v.(int))}
	}