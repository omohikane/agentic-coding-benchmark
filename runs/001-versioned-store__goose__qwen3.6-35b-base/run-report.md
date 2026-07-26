# Run Report

## Summary

- Result: PASS
- Language: Go
- Implementation Files: solution.go
- Final Test Command: go test -v ./...
- Final Test Status: All 16 tests passed (0.002s)
- Implementation Attempts: 1
- Test Executions: 1

## Commands Executed

1. cat README.md
2. cat SPEC.md
3. ls *.go
4. cat go.mod
5. cat solution_test.go
6. write solution.go
7. go test -v ./... | tee test-output.txt
8. gofmt -w .

## Files Changed

- `solution.go` — Created (implementation of versioned store)
- `test-output.txt` — Created (final test output)

## Implementation Summary

Implemented a `Store` type in `solution.go` with the following methods:

- **NewStore()** — Creates an empty key-value store with snapshot support. Initializes internal maps and sets next snapshot ID to 1.
- **Set(key, value)** — Stores or overwrites a key-value pair in the current data map.
- **Get(key)** — Retrieves a value by key; returns `(value, false)` if missing. Supports empty keys and values.
- **Delete(key)** — Removes a key; returns `true` if it existed, `false` otherwise.
- **Snapshot()** — Deep-copies the current data map into a new entry keyed by an incrementing integer ID (starting at 1). Returns the snapshot ID.
- **Restore(snapshotID)** — Validates that the snapshot ID is positive and exists in the snapshots map. If valid, replaces the entire data map with a deep copy of the snapshot's state. Returns `true` on success, `false` for invalid/unknown IDs.

The design ensures:
- Snapshots are independent (deep copies prevent mutation)
- Restoring does not mutate the stored snapshot
- Multiple restores to the same snapshot yield consistent results
- Invalid snapshot IDs (≤ 0 or non-existent) return false without modifying state
- Complex multi-snapshot transitions work correctly

## Test Progression

### Test Run 1

- Command: `go test -v ./...`
- Result: PASS — All 16 tests passed
- Main Failures: None
- Changes Made: N/A (first attempt succeeded)

Test results:
| Test | Status |
|------|--------|
| TestNewStoreIsEmpty | PASS |
| TestSetAndGet | PASS |
| TestSetOverwritesExistingValue | PASS |
| TestEmptyKeyAndEmptyValueAreValid | PASS |
| TestDeleteExistingKey | PASS |
| TestDeleteMissingKey | PASS |
| TestSnapshotIDsIncrease | PASS |
| TestRestoreSnapshot | PASS |
| TestRestoreRemovesKeysCreatedAfterSnapshot | PASS |
| TestRestoreRecoversDeletedKeys | PASS |
| TestSnapshotsAreIndependent | PASS |
| TestRestoringSnapshotDoesNotMutateSnapshot | PASS |
| TestSnapshotAfterRestoreCapturesRestoredState | PASS |
| TestRestoreUnknownSnapshotDoesNothing | PASS |
| TestRestoreInvalidSnapshotIDs | PASS |
| TestComplexStateTransitions | PASS |

## Final Test Output

All 16 tests passed. Complete output is stored in `test-output.txt`.

## Verification

- README modified: No
- SPEC modified: No
- Test files modified: No
- Dependency files modified: No
- Unexpected files created: No (only solution.go, test-output.txt, and run-report.md)
- Accessed parent directories: No

## Self Verification

| Item                                                       | Status   |
| ---------------------------------------------------------- | -------- |
| Read and understood `SPEC.md`                              | Yes      |
| Implemented the required functionality                     | Yes      |
| Executed the final test suite                              | Yes      |
| Saved the final test output to `test-output.txt`           | Yes      |
| Generated `run-report.md`                                  | Yes      |
| Modified any protected files                               | No       |
| Accessed parent directories or other benchmark directories | No       |
| Created unnecessary files                                  | No       |

## Notes

- The SPEC.md was minimal ("reserved for the first task"), so implementation was driven entirely by the test cases in `solution_test.go`.
- All 16 tests passed on the first attempt without any iterations.
- Formatting applied with `gofmt -w .` after implementation.
