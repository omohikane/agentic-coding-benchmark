package versionedstore

// Store is a key-value store that supports snapshots and restores.
type Store struct {
	data      map[string]string
	snapshots map[int]map[string]string
	nextID    int
}

// NewStore creates a new empty Store.
func NewStore() *Store {
	return &Store{
		data:      make(map[string]string),
		snapshots: make(map[int]map[string]string),
		nextID:    1,
	}
}

// Set stores a key-value pair in the store.
func (s *Store) Set(key, value string) {
	s.data[key] = value
}

// Get retrieves the value for a key. Returns ok=false if the key is missing.
func (s *Store) Get(key string) (string, bool) {
	val, ok := s.data[key]
	return val, ok
}

// Delete removes a key from the store. Returns true if the key existed.
func (s *Store) Delete(key string) bool {
	if _, ok := s.data[key]; !ok {
		return false
	}
	delete(s.data, key)
	return true
}

// Snapshot captures the current state of the store and returns a snapshot ID.
// IDs are monotonically increasing positive integers starting from 1.
func (s *Store) Snapshot() int {
	id := s.nextID
	s.nextID++
	// Deep copy the current data into the snapshot
	snap := make(map[string]string, len(s.data))
	for k, v := range s.data {
		snap[k] = v
	}
	s.snapshots[id] = snap
	return id
}

// Restore restores the store to the state captured by the given snapshot ID.
// Returns true if a valid snapshot was found and restored, false otherwise.
func (s *Store) Restore(snapshotID int) bool {
	if snapshotID <= 0 {
		return false
	}
	snap, ok := s.snapshots[snapshotID]
	if !ok {
		return false
	}
	// Replace the entire data map with a deep copy of the snapshot
	s.data = make(map[string]string, len(snap))
	for k, v := range snap {
		s.data[k] = v
	}
	return true
}
