package todo

import (
	"sync"

	"github.com/google/uuid"
)

type Service struct {
	mu    sync.RWMutex
	items map[string]Item
}

func NewService() *Service {
	return &Service{
		items: make(map[string]Item),
	}
}

func (s *Service) List() []Item {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]Item, 0, len(s.items))
	for _, it := range s.items {
		out = append(out, it)
	}
	return out
}

func (s *Service) Create(title string) Item {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := uuid.NewString()
	it := Item{ID: id, Title: title}
	s.items[id] = it
	return it
}

func (s *Service) Delete(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.items[id]; !ok {
		return false
	}
	delete(s.items, id)
	return true
}
