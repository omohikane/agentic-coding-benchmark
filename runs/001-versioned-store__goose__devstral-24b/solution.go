package versionedstore

import "sync"

// Store is a key-value store that supports snapshot and restore.
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

// Set stores a key-value pair in the store.
func (s *Store) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

// Get retrieves the value for a key. Returns false if the key is missing.
func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.data[key]
	return val, ok
}

// Delete removes a key from the store. Returns true if the key existed.
func (s *Store) Delete(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.data[key]; !ok {
		return false
	}
	delete(s.data, key)
	return true
}

// Snapshot captures the current state and returns a unique snapshot ID.
func (s *Store) Snapshot() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := s.nextID
	s.nextID++

	// Deep copy current data for this snapshot
	cp := make(map[string]string, len(s.data))
	for k, v := range s.data {
		cp[k] = v
	}
	s.snapshots[id] = cp

	return id
}

// Restore restores the store to the state captured by the given snapshot ID.
// Returns true if the snapshot was found and restored, false otherwise.
func (s *Store) Restore(snapshotID int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	cp, ok := s.snapshots[snapshotID]
	if !ok || snapshotID <= 0 {
		return false
	}

	// Replace current data with the snapshot's state
	s.data = make(map[string]string, len(cp))
	for k, v := range cp {
		s.data[k] = v
	}

	return true
}
