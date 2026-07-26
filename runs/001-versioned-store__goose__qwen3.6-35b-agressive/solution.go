package versionedstore

import "sync"

// Snapshot represents a point-in-time copy of the store's state.
type Snapshot struct {
	ID    int
	Items map[string]string
}

// Store is a versioned key-value store that supports snapshots and restores.
type Store struct {
	mu        sync.RWMutex
	items     map[string]string
	snapshots map[int]Snapshot
	nextID    int
}

// NewStore creates a new empty Store.
func NewStore() *Store {
	return &Store{
		items:     make(map[string]string),
		snapshots: make(map[int]Snapshot),
		nextID:    1,
	}
}

// Set sets the value for the given key.
func (s *Store) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[key] = value
}

// Get retrieves the value for the given key. Returns the value and true if found,
// or an empty string and false if not found.
func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.items[key]
	return val, ok
}

// Delete removes the key from the store. Returns true if the key existed and was deleted.
func (s *Store) Delete(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.items[key]
	if ok {
		delete(s.items, key)
	}
	return ok
}

// Snapshot takes a point-in-time snapshot of the store and returns its ID.
func (s *Store) Snapshot() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := s.nextID
	s.nextID++

	// Deep copy the items map for this snapshot
	itemsCopy := make(map[string]string, len(s.items))
	for k, v := range s.items {
		itemsCopy[k] = v
	}

	s.snapshots[id] = Snapshot{
		ID:    id,
		Items: itemsCopy,
	}

	return id
}

// Restore restores the store to the state captured by the snapshot with the given ID.
// Returns true if a valid snapshot was found and restored, false otherwise.
func (s *Store) Restore(snapshotID int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	snaps, ok := s.snapshots[snapshotID]
	if !ok || snapshotID <= 0 {
		return false
	}

	// Replace the current state with the snapshot's state
	s.items = make(map[string]string, len(snaps.Items))
	for k, v := range snaps.Items {
		s.items[k] = v
	}

	return true
}
