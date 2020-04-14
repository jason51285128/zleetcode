package zleetcode

import (
	"container/list"
	"time"
)

//实现一个带过期时间的LRU算法
//lru过期采用惰性机制
type lruValue struct {
	key    interface{}
	value  interface{}
	expire int64
}

type lru struct {
	cachelist *list.List
	cachemap  map[interface{}]*list.Element
	capacity  int
}

func newLru(c int) *lru {
	return &lru{
		capacity:  c,
		cachelist: list.New(),
		cachemap:  make(map[interface{}]*list.Element),
	}
}

func (l *lru) get(key interface{}) interface{} {
	if e, ok := l.cachemap[key]; ok {
		if e.Value.(*lruValue).expire > time.Now().Unix() {
			l.cachelist.MoveToFront(e)
			return e.Value.(*lruValue).value
		}
		l.cachelist.Remove(e)
		delete(l.cachemap, key)
	}
	return nil
}

func (l *lru) set(key interface{}, value interface{}, expire int64) {
	if e, ok := l.cachemap[key]; ok {
		if expire > time.Now().Unix() {
			e.Value.(*lruValue).expire = expire
			e.Value.(*lruValue).value = value
			return
		}
		l.cachelist.Remove(e)
		delete(l.cachemap, key)
		return
	}
	if expire < time.Now().Unix() {
		return
	}
	if l.cachelist.Len() >= l.capacity {
		oe := l.cachelist.Remove(l.cachelist.Back())
		delete(l.cachemap, oe.(*lruValue).key)

	}
	ne := l.cachelist.PushFront(&lruValue{key: key, expire: expire, value: value})
	l.cachemap[key] = ne
}
