type MyHashSet struct {
	data []bool
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{make([]bool, 1000001)}
}

func (this *MyHashSet) Add(key int) {
	this.data[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.data[key] = false
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.data[key]
}
