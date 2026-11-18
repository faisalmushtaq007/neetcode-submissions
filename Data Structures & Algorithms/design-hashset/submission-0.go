type MyHashSet struct {
	Keys []int
}

func Constructor() MyHashSet {
	return MyHashSet{
		Keys: []int{},
	}
}

func (this *MyHashSet) Add(key int) {
    if !this.Contains(key) {
		this.Keys = append(this.Keys, key)
	}
}

func (this *MyHashSet) Remove(key int) {
   for i, k := range this.Keys {
		if k == key {
			this.Keys[i] = this.Keys[len(this.Keys)-1]
			this.Keys = this.Keys[:len(this.Keys)-1]
			return
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
    for _, value := range this.Keys {
		if key == value {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 