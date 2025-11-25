package main

import "sort"

type pair struct {
	value string
	timestamp int
}

type TimeMap struct {
	mp map[string][]pair
}


func Constructor() TimeMap {
  return TimeMap{mp: map[string][]pair{}}
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
  this.mp[key] = append(this.mp[key], pair{value, timestamp})
}


func (this *TimeMap) Get(key string, timestamp int) string {
  list := this.mp[key]
	idx := sort.Search(len(list), func(i int) bool {
		return list[i].timestamp > timestamp
	})
	if idx > 0 {
		return list[idx-1].value
	}
	return ""
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
