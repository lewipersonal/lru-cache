package lewilru

import "container/list"
import "errors"
import "fmt"

type LRUCache struct {
	items  map[string]*list.Element
	lookup *list.List
	size   int
}

type element struct {
	key string
	value int
}

func New(size int) (*LRUCache, error) {
	if size <= 0 {
		return nil, errors.New("LRU Cache must have positive size")
	}

	lru := &LRUCache{
		items:  make(map[string]*list.Element),
		lookup: list.New(),
		size:   size,
	}

	return lru, nil
}

func (v *LRUCache) Set(key string, value int) {
	// if it already exists then move it to latest
	fmt.Println("Setting", key)
	val, ok := v.items[key]
	if ok {
		if val.Value.(*element).value != value {
			fmt.Println("replacing", val.Value.(*element).value, "with", value)
			val.Value.(*element).value = value
		}
		v.lookup.MoveToFront(val)
		fmt.Println("Cache now", v.GetAll())
		return
	}

	// otherwise set it
	e := v.lookup.PushFront(&element{ key, value })
	v.items[key] = e

	// check size and evict if necessary
	if v.lookup.Len() > v.size {
		v.evict()
	}

	fmt.Println("Cache now", v.GetAll())
	return
}

func (v *LRUCache) Get(key string) (int, error) {
	fmt.Println("Getting", key)

	val, ok := v.items[key]
	if ok {
		v.lookup.MoveToFront(val)
		fmt.Println("Key", key, "has value", val.Value.(*element).value)
		return val.Value.(*element).value, nil
	}

	fmt.Println("Key", key, "not present")
	return 0, errors.New("Key not present")
}

func (v *LRUCache) GetAll() []element {
	items := make([]element, v.lookup.Len())
	node := v.lookup.Front()
	for i := 0; i < v.lookup.Len(); i++ {
		items[i] = *node.Value.(*element)
		node = node.Next()
	}
	return items
}

func (v *LRUCache) evict() {
	e := v.lookup.Back()
	delete(v.items, (e.Value.(*element).key))
	v.lookup.Remove(e)
	fmt.Println("evicted", e.Value.(*element).key)
	return
}