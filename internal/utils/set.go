package utils

import "strings"

type EmptyElement struct{}
type Set struct {
	container map[string]EmptyElement
}

func NewSet() *Set {
	m := make(map[string]EmptyElement, 0)
	return &Set{
		container: m,
	}
}

func (s *Set) Add(val string) {
	for key, _ := range s.container {
		if strings.Contains(val, key) && len(val) > len(key) {
			s.Remove(key)
		}
	}
	s.container[val] = EmptyElement{}
}

func (s *Set) Remove(val string) {
	delete(s.container, val)
}

func (s *Set) Contains(val string) bool {
	_, ok := s.container[val]
	return ok
}

func (s *Set) ConvertToArray() []string {
	var result []string
	for key, _ := range s.container {
		result = append(result, key)
	}
	return result
}
