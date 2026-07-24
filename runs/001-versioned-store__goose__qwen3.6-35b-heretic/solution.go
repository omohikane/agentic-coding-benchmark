package versionedstore

import "sync"

type Store struct {
	mu        sync.RWMutex
	data      map[string]string
	snapshots map[int]map[string]string
	nextID    int
}

func NewStore() *Store {
	return &Store{
		data:      make(map[string]string),
		snapshots: make(map[int]map[string]string),
		nextID:    1,
	}
}

func (s *Store) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.data[key]
	return v, ok
}

func (s *Store) Delete(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.data[key]
	if !exists {
		return false
	}
	delete(s.data, key)
	return true
}

func (s *Store) Snapshot() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := s.nextID
	s.nextID++
	cp := make(map[string]string, len(s.data))
	for k, v := range s.data {
		cp[k] = v
	}
	s.snapshots[id] = cp
	return id
}

func (s *Store) Restore(snapshotID int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	cp, ok := s.snapshots[snapshotID]
	if !ok || snapshotID <= 0 {
		return false
	}
	s.data = make(map[string]string, len(cp))
	for k, v := range cp {
		s.data[k] = v
	}
	return true
}
