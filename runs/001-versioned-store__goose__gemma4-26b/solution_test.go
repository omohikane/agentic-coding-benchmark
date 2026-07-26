package versionedstore

import "testing"

func requireValue(t *testing.T, store *Store, key, expected string) {
	t.Helper()

	got, ok := store.Get(key)
	if !ok {
		t.Fatalf("Get(%q): key not found; want %q", key, expected)
	}

	if got != expected {
		t.Fatalf("Get(%q) = %q; want %q", key, got, expected)
	}
}

func requireMissing(t *testing.T, store *Store, key string) {
	t.Helper()

	got, ok := store.Get(key)
	if ok {
		t.Fatalf("Get(%q) = (%q, true); want key to be missing", key, got)
	}
}

func TestNewStoreIsEmpty(t *testing.T) {
	store := NewStore()

	requireMissing(t, store, "missing")
}

func TestSetAndGet(t *testing.T) {
	store := NewStore()

	store.Set("language", "go")
	requireValue(t, store, "language", "go")
}

func TestSetOverwritesExistingValue(t *testing.T) {
	store := NewStore()

	store.Set("version", "1")
	store.Set("version", "2")

	requireValue(t, store, "version", "2")
}

func TestEmptyKeyAndEmptyValueAreValid(t *testing.T) {
	store := NewStore()

	store.Set("", "")
	requireValue(t, store, "", "")

	store.Set("empty", "")
	requireValue(t, store, "empty", "")
}

func TestDeleteExistingKey(t *testing.T) {
	store := NewStore()

	store.Set("temporary", "value")

	if deleted := store.Delete("temporary"); !deleted {
		t.Fatal("Delete existing key returned false; want true")
	}

	requireMissing(t, store, "temporary")
}

func TestDeleteMissingKey(t *testing.T) {
	store := NewStore()

	if deleted := store.Delete("missing"); deleted {
		t.Fatal("Delete missing key returned true; want false")
	}
}

func TestSnapshotIDsIncrease(t *testing.T) {
	store := NewStore()

	first := store.Snapshot()
	second := store.Snapshot()
	third := store.Snapshot()

	if first <= 0 {
		t.Fatalf("first Snapshot() = %d; want a positive ID", first)
	}

	if second <= first {
		t.Fatalf("second Snapshot() = %d; want greater than first ID %d", second, first)
	}

	if third <= second {
		t.Fatalf("third Snapshot() = %d; want greater than second ID %d", third, second)
	}
}

func TestRestoreSnapshot(t *testing.T) {
	store := NewStore()

	store.Set("status", "before")
	snapshotID := store.Snapshot()

	store.Set("status", "after")
	requireValue(t, store, "status", "after")

	if restored := store.Restore(snapshotID); !restored {
		t.Fatalf("Restore(%d) returned false; want true", snapshotID)
	}

	requireValue(t, store, "status", "before")
}

func TestRestoreRemovesKeysCreatedAfterSnapshot(t *testing.T) {
	store := NewStore()

	store.Set("existing", "value")
	snapshotID := store.Snapshot()

	store.Set("new", "later")

	if restored := store.Restore(snapshotID); !restored {
		t.Fatalf("Restore(%d) returned false; want true", snapshotID)
	}

	requireValue(t, store, "existing", "value")
	requireMissing(t, store, "new")
}

func TestRestoreRecoversDeletedKeys(t *testing.T) {
	store := NewStore()

	store.Set("recoverable", "original")
	snapshotID := store.Snapshot()

	store.Delete("recoverable")
	requireMissing(t, store, "recoverable")

	if restored := store.Restore(snapshotID); !restored {
		t.Fatalf("Restore(%d) returned false; want true", snapshotID)
	}

	requireValue(t, store, "recoverable", "original")
}

func TestSnapshotsAreIndependent(t *testing.T) {
	store := NewStore()

	store.Set("value", "one")
	first := store.Snapshot()

	store.Set("value", "two")
	second := store.Snapshot()

	store.Set("value", "three")

	if restored := store.Restore(first); !restored {
		t.Fatalf("Restore(%d) returned false; want true", first)
	}
	requireValue(t, store, "value", "one")

	if restored := store.Restore(second); !restored {
		t.Fatalf("Restore(%d) returned false; want true", second)
	}
	requireValue(t, store, "value", "two")
}

func TestRestoringSnapshotDoesNotMutateSnapshot(t *testing.T) {
	store := NewStore()

	store.Set("counter", "1")
	snapshotID := store.Snapshot()

	if restored := store.Restore(snapshotID); !restored {
		t.Fatalf("first Restore(%d) returned false; want true", snapshotID)
	}

	store.Set("counter", "999")

	if restored := store.Restore(snapshotID); !restored {
		t.Fatalf("second Restore(%d) returned false; want true", snapshotID)
	}

	requireValue(t, store, "counter", "1")
}

func TestSnapshotAfterRestoreCapturesRestoredState(t *testing.T) {
	store := NewStore()

	store.Set("branch", "original")
	original := store.Snapshot()

	store.Set("branch", "changed")
	store.Set("extra", "present")

	if restored := store.Restore(original); !restored {
		t.Fatalf("Restore(%d) returned false; want true", original)
	}

	restoredState := store.Snapshot()

	store.Set("branch", "mutated-again")
	store.Set("extra", "new")

	if restored := store.Restore(restoredState); !restored {
		t.Fatalf("Restore(%d) returned false; want true", restoredState)
	}

	requireValue(t, store, "branch", "original")
	requireMissing(t, store, "extra")
}

func TestRestoreUnknownSnapshotDoesNothing(t *testing.T) {
	store := NewStore()

	store.Set("stable", "value")

	if restored := store.Restore(999999); restored {
		t.Fatal("Restore with unknown snapshot ID returned true; want false")
	}

	requireValue(t, store, "stable", "value")
}

func TestRestoreInvalidSnapshotIDs(t *testing.T) {
	store := NewStore()

	store.Set("stable", "value")

	for _, snapshotID := range []int{0, -1, -100} {
		if restored := store.Restore(snapshotID); restored {
			t.Errorf("Restore(%d) returned true; want false", snapshotID)
		}

		requireValue(t, store, "stable", "value")
	}
}

func TestComplexStateTransitions(t *testing.T) {
	store := NewStore()

	store.Set("a", "1")
	store.Set("b", "1")
	first := store.Snapshot()

	store.Set("a", "2")
	store.Delete("b")
	store.Set("c", "2")
	second := store.Snapshot()

	store.Set("a", "3")
	store.Set("b", "3")
	store.Delete("c")

	if restored := store.Restore(first); !restored {
		t.Fatalf("Restore(%d) returned false; want true", first)
	}

	requireValue(t, store, "a", "1")
	requireValue(t, store, "b", "1")
	requireMissing(t, store, "c")

	store.Set("a", "from-first")
	store.Delete("b")

	if restored := store.Restore(second); !restored {
		t.Fatalf("Restore(%d) returned false; want true", second)
	}

	requireValue(t, store, "a", "2")
	requireMissing(t, store, "b")
	requireValue(t, store, "c", "2")

	if restored := store.Restore(first); !restored {
		t.Fatalf("second Restore(%d) returned false; want true", first)
	}

	requireValue(t, store, "a", "1")
	requireValue(t, store, "b", "1")
	requireMissing(t, store, "c")
}
