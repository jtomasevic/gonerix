package collections

type nodeWithKey[K comparable, V comparable] struct {
	Key   K
	Value V
}

// lruCache Least Recently Used (LRU) is a common caching strategy.
// It defines the policy to evict elements from the cache to make room for new elements when the cache is full,
// meaning it discards the least recently used items first.
type lruCache[K comparable, V comparable] struct {
	capacity  int
	cache     linkedList[nodeWithKey[K, V]]
	cacheHash Dictionary[K, *Node[nodeWithKey[K, V]]]
	headKey   K
}

func LRUCache[K comparable, V comparable](capacity int) lruCache[K, V] {
	return lruCache[K, V]{
		capacity:  capacity,
		cache:     LinkedList[nodeWithKey[K, V]](),
		cacheHash: Dictionary[K, *Node[nodeWithKey[K, V]]]{},
	}
}

// Set If key doesn't exist add new item to cache. If cache is full, remove last used.
//  If key exist set new value in cache, and set item to be first used (last removing from current cache state)
//  O(n) = 1
func (lru *lruCache[K, V]) Set(key K, value V) {
	node, ok := lru.cacheHash[key]
	if ok {
		node.Value.Value = value
		lru.cache.RemoveNode(node)
		lru.cache.AddAfter(lru.cache.tail, node.Value)
	} else {
		length := lru.cacheHash.Length()
		if lru.capacity == length {
			lru.cache.RemoveNode(lru.cache.head)
			lru.cacheHash.Remove(lru.cache.head.Value.Key)
		}
		addedNode := lru.cache.AddAfter(lru.cache.tail, nodeWithKey[K, V]{
			Key:   key,
			Value: value,
		})
		lru.cacheHash[key] = addedNode
	}
}

// Get If key exist return value from cache,
//  set item to be first used (last removing from current cache state)
//  O(n) = 1
func (lru *lruCache[K, V]) Get(key K) *V {
	node, ok := lru.cacheHash[key]
	if ok {
		lru.cache.RemoveNode(node)
		lru.cache.AddAfter(lru.cache.tail, node.Value)
		return &node.Value.Value
	}
	return nil
}

// GetCacheValues Return slice of values from cache in order from last visit to first visit.
//  Be careful with this operation because O(n) = n
func (lru *lruCache[K, V]) GetCacheValues() []V {
	var result []V
	for _, v := range lru.cache.Values() {
		result = append(result, v.Value)
	}
	return result
}
