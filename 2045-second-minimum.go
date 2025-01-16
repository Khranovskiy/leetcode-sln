package main

import (
	"container/heap"
	"fmt"
)

type Edge struct {
	node, time int
}

type MinHeap []Edge

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].time < h[j].time }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Edge))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func secondMinimum(n int, edges [][]int, time int, change int) int {
	// Построение графа
	graph := make([][]int, n+1)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// Массив для минимального и второго минимального времени
	times := make([][2]int, n+1)
	for i := range times {
		times[i] = [2]int{1 << 31, 1 << 31} // Используем большой int (эквивалент бесконечности)
	}
	times[1][0] = 0

	// MinHeap для реализации BFS
	h := &MinHeap{}
	heap.Init(h)
	heap.Push(h, Edge{node: 1, time: 0})

	for h.Len() > 0 {
		e := heap.Pop(h).(Edge)
		node, currTime := e.node, e.time

		for _, neighbor := range graph[node] {
			newTime := currTime

			// Проверка сигнала светофора
			if (currTime/change)%2 == 1 {
				newTime += change - (currTime % change)
			}
			newTime += time

			if newTime < times[neighbor][0] {
				times[neighbor][1] = times[neighbor][0]
				times[neighbor][0] = newTime
				heap.Push(h, Edge{node: neighbor, time: newTime})
			} else if newTime > times[neighbor][0] && newTime < times[neighbor][1] {
				times[neighbor][1] = newTime
				heap.Push(h, Edge{node: neighbor, time: newTime})
			}
		}
	}

	// Возвращаем второе минимальное время для n
	return times[n][1]
}

func main() {
	n := 5
	edges := [][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}, {4, 5}}
	time := 3
	change := 5
	fmt.Println(secondMinimum(n, edges, time, change))
}
