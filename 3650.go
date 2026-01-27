package main

import (
	"container/heap"
	"math"
)

type Number interface {
	int | int64 | float64 | float32
}

type edge[T Number] struct {
	to     int
	weight T
}

type item[T Number] struct {
	node int
	val  T
}

type priorityQueue[T Number] []*item[T]

func (pq priorityQueue[T]) Len() int           { return len(pq) }
func (pq priorityQueue[T]) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQueue[T]) Less(i, j int) bool { return pq[i].val < pq[j].val }
func (pq *priorityQueue[T]) Push(x any)         { *pq = append(*pq, x.(*item[T])) }
func (pq *priorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

// Dijkstra 结构体
type DijkstraSolver[T Number] struct {
	n   int
	g   [][]edge[T]
	dis []T
}

// 构造函数
func NewDijkstraSolver[T Number](n int) *DijkstraSolver[T] {
	return &DijkstraSolver[T]{
		n: n,
		g: make([][]edge[T], n),
	}
}

// 添加有向边 (单向)
// 如果是无向图，请调用两次：AddEdge(u, v, w); AddEdge(v, u, w);
func (d *DijkstraSolver[T]) AddEdge(u, v int, weight T) {
	d.g[u] = append(d.g[u], edge[T]{to: v, weight: weight})
}

func (ds *DijkstraSolver[T]) ShortestPath(s, t int) T {
	var maxT T
	maxT = math.MaxInt

	ds.dis = make([]T, ds.n)
	for i := range ds.dis {
		ds.dis[i] = maxT
	}

	pq := &priorityQueue[T]{}
	heap.Init(pq)

	ds.dis[s] = 0
	heap.Push(pq, &item[T]{val: 0, node: s})

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*item[T])
		d, u := curr.val, curr.node

		if d > ds.dis[u] {
			continue
		}

		if u == t {
			return d
		}

		for _, e := range ds.g[u] {
			v := e.to
			weight := e.weight

			if ds.dis[v] > d+weight {
				ds.dis[v] = d + weight
				heap.Push(pq, &item[T]{val: ds.dis[v], node: v})
			}
		}
	}
	return ds.dis[t]
}

func minCost(n int, edges [][]int) int {

	solver := NewDijkstraSolver[int](n)

	for i := range edges {
		solver.AddEdge(edges[i][0], edges[i][1], edges[i][2])
		solver.AddEdge(edges[i][1], edges[i][0], 2*edges[i][2])
	}

	dist := solver.ShortestPath(0, n-1)
	if dist == math.MaxInt {
		return -1
	}
	return dist
}
