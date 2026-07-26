package versionedstore

// Store provides a key-value store with snapshot and restore capabilities.
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

// Get retrieves the value for a given key. Returns false if the key is missing.
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
// Snapshot IDs are positive integers that strictly increase with each call.
func (s *Store) Snapshot() int {
	id := s.nextID
	s.nextID++

	// Deep copy the current data for this snapshot
	copy := make(map[string]string, len(s.data))
	for k, v := range s.data {
		copy[k] = v
	}
	s.snapshots[id] = copy

	return id
}

// Restore restores the store to the state captured in the given snapshot.
// Returns true if the snapshot was found and restored successfully.
// Restoring a snapshot removes keys that were added after the snapshot,
// and recovers keys that existed at snapshot time but were deleted later.
func (s *Store) Restore(snapshotID int) bool {
	if snapshotID <= 0 {
		return false
	}

	snap, ok := s.snapshots[snapshotID]
	if !ok {
		return false
	}

	// Replace the current data with a deep copy of the snapshot state
	newData := make(map[string]string, len(snap))
	for k, v := range snap {
		newData[k] = v
	}
	s.data = newData

	return true
}
