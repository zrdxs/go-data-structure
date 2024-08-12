package lcucache

import "container/list"

type KeyPair struct {
	key   int
	value int
}

type LRUCache struct {
	capacity int
	list     *list.List
	elements map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	elements := make(map[int]*list.Element)
	return LRUCache{
		capacity: capacity,
		list:     list.New(),
		elements: elements,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.elements[key]; ok {
		value := node.Value.(list.Element).Value.(KeyPair).value
		this.list.MoveToFront(node)
		return value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.elements[key]; ok {
		this.list.MoveToFront(node)
		actualValue := node.Value.(list.Element).Value.(KeyPair).value
		node.Value.(*list.Element).Value = KeyPair{key: key, value: value + actualValue}
	} else {
		if this.list.Len() == this.capacity {
			//the element key is the same from pointer key at list
			elemIdx := this.list.Back().Value.(list.Element).Value.(KeyPair).key
			delete(this.elements, elemIdx)
			this.list.Remove(this.list.Back())
		}
	}

	node := list.Element{
		Value: KeyPair{
			key:   key,
			value: value,
		},
	}

	pointer := this.list.PushFront(node)
	this.elements[key] = pointer
}
