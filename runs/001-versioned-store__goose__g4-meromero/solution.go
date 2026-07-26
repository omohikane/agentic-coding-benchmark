package versionedstore

// Store provides a key-value store with snapshot and restore capabilities.
type Store struct {
	current   map[string]string
	snapshots map[int]map[string]string
	nextID    int
}

// NewStore creates an empty Store.
func NewStore() *Store {
	return &Store{
		current:   make(map[string]string),
		snapshots: make(map[int]map[string]string),
		nextID:    1,
	}
}

// Set stores a key-value pair in the store.
func (s *Store) Set(key, value string) {
	s.current[key] = value
}

// Get retrieves a value by key. Returns false if the key does not exist.
func (s *Store) Get(key string) (string, bool) {
	val, ok := s.current[key]
	return val, ok
}

// Delete removes a key from the store. Returns true if the key existed.
func (s *Store) Delete(key string) bool {
	if _, ok := s.current[key]; !ok {
		return false
	}
	delete(s.current, key)
	return true
}

// Snapshot captures the current state and returns a snapshot ID.
func (s *Store) Snapshot() int {
	id := s.nextID
	s.nextID++
	cp := make(map[string]string, len(s.current))
	for k, v := range s.current {
		cp[k] = v
	}
	s.snapshots[id] = cp
	return id
}

// Restore restores the store to the state captured at the given snapshot ID.
// Returns false if the snapshot ID is invalid (<= 0 or not found).
func (s *Store) Restore(snapshotID int) bool {
	if snapshotID <= 0 {
		return false
	}
	snap, ok := s.snapshots[snapshotID]
	if !ok {
		return false
	}
	cp := make(map[string]string, len(snap))
	for k, v := range snap {
		cp[k] = v
	}
	s.current = cp
	return true
}
