package collections



type Dictionary[K comparable, V comparable] map[K]V

func (dic *Dictionary[K, V]) Remove(key K) (bool, V) {
	current:= *dic
	for k, value :=range current {
		if k == key {
			delete(current, k)
			dic = &current
			return true, value
		}
	}
	var empty V
	return false, empty
}

func (dic *Dictionary[K, V]) Length() int {
	return len(*dic)
}