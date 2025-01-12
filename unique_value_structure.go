// https://algorithmics-blog.github.io/blog/unique_value_structure/
package main

import "math/rand"

type Store struct {
	indexesMap map[int]int
	values     []int
}

func New() *Store {
	return &Store{
		indexesMap: make(map[int]int),
		values:     make([]int, 0),
	}
}

func (s *Store) Insert(value int) {
	_, exists := s.indexesMap[value]
	if exists {
		return
	}
	s.values = append(s.values, value)
	s.indexesMap[value] = len(s.values) - 1
}

func (s *Store) Remove(value int) {
	index, exists := s.indexesMap[value]
	if !exists {
		return
	}
	// if we deletes element from value array in the middle
	// we should update indexes of right side elements in indexesMap

	// but we can only move the last element to deleted position
	// without updating

	last := s.values[len(s.values)-1]
	s.values[index] = last
	s.indexesMap[last] = index

	// remove last element
	s.values = s.values[:len(s.values)-1]
	delete(s.indexesMap, value)
}

func (s *Store) GetRandom() int {
	index := rand.Intn(len(s.values))
	return s.values[index]
}
