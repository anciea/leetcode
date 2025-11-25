package main

// import "sort"

// type RangeFreqQuery struct {
// 	vis map[int][]int
// }


// func Constructor(arr []int) RangeFreqQuery {
// 	vis := make(map[int][]int)
// 	for i,x := range arr {
// 		if _,ok := vis[x]; !ok {
// 			vis[x] = make([]int,0)
// 		}
// 		vis[x] = append(vis[x], i)
// 	}
// 	return RangeFreqQuery{vis : vis}
// }


// func (this *RangeFreqQuery) Query(left int, right int, value int) int {
//   list := this.vis[value]
// 	l := sort.SearchInts(list, left)
// 	r := sort.SearchInts(list, right+1)
// 	return r - l
// }


// /**
//  * Your RangeFreqQuery object will be instantiated and called as such:
//  * obj := Constructor(arr);
//  * param_1 := obj.Query(left,right,value);
//  */
