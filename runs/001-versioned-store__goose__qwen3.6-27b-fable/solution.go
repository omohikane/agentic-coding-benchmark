package versionedstore

import "sync"

// Store is a key-value store that supports snapshotting and restoring.
type Store struct {
	mu        sync.RWMutex
	data      map[string]string
	snapshots map[int]map[string]string
	nextID    int
}

// NewStore creates a new empty versioned store.
func NewStore() *Store {
	return &Store{
		data:      make(map[string]string),
		snapshots: make(map[int]map[string]string),
		nextID:    1,
	}
}

// Set sets the value for a key in the current store.
func (s *Store) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

// Get retrieves the value for a key. Returns ("", false) if the key is missing.
func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.data[key]
	return val, ok
}

// Delete removes a key from the current store. Returns true if the key existed.
func (s *Store) Delete(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.data[key]
	if ok {
		delete(s.data, key)
	}
	return ok
}

// Snapshot takes a snapshot of the current store state and returns a unique ID.
func (s *Store) Snapshot() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	id := s.nextID
	s.nextID++

	// Deep copy the current data for this snapshot
	snap := make(map[string]string, len(s.data))
	for k, v := range s.data {
		snap[k] = v
	}

	s.snapshots[id] = snap
	return id
}

// Restore restores the store to a previous snapshot state. Returns true if found.
func (s *Store) Restore(snapshotID int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	snap, ok := s.snapshots[snapshotID]
	if !ok || snapshotID <= 0 {
		return false
	}

	// Replace the current data with a deep copy of the snapshot
	s.data = make(map[string]string, len(snap))
	for k, v := range snap {
		s.data[k] = v
	}
	return true
}
