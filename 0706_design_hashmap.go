type MyHashMap struct {
	vals []int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	m := MyHashMap{make([]int, 100)}
	for i := 0; i < 100; i++ {
		m.vals[i] = -1
	}
	return m
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	if key >= len(this.vals) {
		prev_len := len(this.vals)
		s2 := make([]int, key+1-prev_len)
		this.vals = append(this.vals, s2...)
		for i := prev_len; i < len(this.vals); i++ {
			this.vals[i] = -1
		}
	}
	this.vals[key] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	if key > len(this.vals) {
		return -1
	}
	return this.vals[key]
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	if key > len(this.vals) {
		return
	}
	this.vals[key] = -1
}
