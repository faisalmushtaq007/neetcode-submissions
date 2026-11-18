type MyHashMap struct {
    Values []int
}

func Constructor() MyHashMap {
    return MyHashMap{
        Values: make([]int, 0),
    }
}

func (this *MyHashMap) Put(key int, value int) {
    if key >= len(this.Values) {
        newSlice := make([]int, key+1)
        copy(newSlice, this.Values)
        this.Values = newSlice
    }
    this.Values[key] = value + 1
}

func (this *MyHashMap) Get(key int) int {
    if key >= len(this.Values) || this.Values[key] == 0 {
        return -1
    }
    return this.Values[key] - 1
}

func (this *MyHashMap) Remove(key int) {
    if key < len(this.Values) {
        this.Values[key] = 0
    }
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */