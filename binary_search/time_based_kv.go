package binarysearch

import "sort"

// Design a time-based key-value data structure that can store multiple values for
// the same key at different time stamps and retrieve the key's value at a certain timestamp.
// Implement the TimeMap class:
// TimeMap() Initializes the object of the data structure.
// void set(String key, String value, int timestamp) Stores the key
// key with the value value at the given time timestamp.
// String get(String key, int timestamp) Returns a value such that
// set was called previously, with timestamp_prev <= timestamp. If there are multiple
// such values, it returns the value associated with the largest timestamp_prev.
// If there are no values, it returns "".

type tuple struct {
	timestamp int
	value     string
}
type TimeMap struct {
	values map[string][]*tuple
}

func Constructor() TimeMap {
	return TimeMap{
		values: make(map[string][]*tuple),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if key == "" {
		return
	}

	t := &tuple{

		timestamp,
		value,
	}
	_, ok := this.values[key]
	if !ok {
		this.values[key] = []*tuple{t}
		return
	}
	this.values[key] = append(this.values[key], t)
}

func (this *TimeMap) Get(knordey string, timestamp int) string {
	val, ok := this.values[knordey]
	if !ok {
		return ""
	}
	i := sort.Search(len(val), func(i int) bool { return val[i].timestamp < timestamp })
	return val[i-1].value
}

func search(list []*tuple, timestamp int) string {
	res := ""
	left, right := 0, len(list)-1
	for left <= right {
		mid := (left + right) / 2
		if list[mid].timestamp <= timestamp {
			res = list[mid].value
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}
func (t *TimeMap) Get2(key string, timestamp int) string {
	values, ok := t.values[key]
	if !ok {
		return ``
	}

	var prev string
	left, right := 0, len(values)-1

	for left <= right {
		mid := left + (right-left)/2

		switch {
		case values[mid].timestamp == timestamp:
			return values[mid].value
		case values[mid].timestamp < timestamp:
			left = mid + 1
			prev = values[mid].value
		default:
			right = mid - 1
		}
	}

	return prev
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
