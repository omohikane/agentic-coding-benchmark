# Run Report

## Summary

- Result: PASS
- Language: Go
- Implementation Files: solution.go
- Final Test Command: `go test -v ./...`
- Final Test Status: All 16 tests passed
- Implementation Attempts: 1
- Test Executions: 1

## Commands Executed

1. cat README.md
2. cat SPEC.md
3. cat go.mod
4. cat solution_test.go
5. write solution.go (implement versioned store)
6. go test -v ./...
7. gofmt -w .
8. tee test-output.txt

## Files Changed

- `solution.go` — created (new implementation of the versioned key-value store with snapshot/restore support)

## Implementation Summary

Implemented a thread-safe versioned key-value store (`Store`) in Go with the following operations:

- **NewStore()** — creates an empty store with initialized data map, snapshots map, and snapshot ID counter starting at 1.
- **Set(key, value)** — stores or overwrites a key-value pair (protected by write lock).
- **Get(key)** — retrieves a value; returns `(value, false)` if the key is missing.
- **Delete(key)** — removes a key; returns `true` if it existed, `false` otherwise.
- **Snapshot()** — captures a deep copy of the current state and assigns a monotonically increasing ID starting from 1. Returns the snapshot ID.
- **Restore(snapshotID)** — replaces the current store state with a deep copy of the saved snapshot. Returns `true` on success, `false` for unknown or invalid (≤0) snapshot IDs. Keys present after the snapshot but absent in it are removed; keys deleted after the snapshot are restored.

The implementation uses `sync.RWMutex` for concurrency safety and performs deep copies during both Snapshot and Restore to ensure snapshots remain independent of subsequent mutations.

## Test Progression

### Test Run 1

- Command: `go test -v ./...`
- Result: PASS (all 16 tests passed)
- Main Failures: None
- Changes Made: N/A — first attempt succeeded

| # | Test Name                              | Status |
|---|----------------------------------------|--------|
| 1 | TestNewStoreIsEmpty                    | PASS   |
| 2 | TestSetAndGet                          | PASS   |
| 3 | TestSetOverwritesExistingValue         | PASS   |
| 4 | TestEmptyKeyAndEmptyValueAreValid      | PASS   |
| 5 | TestDeleteExistingKey                  | PASS   |
| 6 | TestDeleteMissingKey                   | PASS   |
| 7 | TestSnapshotIDsIncrease                | PASS   |
| 8 | TestRestoreSnapshot                    | PASS   |
| 9 | TestRestoreRemovesKeysCreatedAfterSnapshot | PASS |
| 10| TestRestoreRecoversDeletedKeys         | PASS   |
| 11| TestSnapshotsAreIndependent            | PASS   |
| 12| TestRestoringSnapshotDoesNotMutateSnapshot | PASS |
| 13| TestSnapshotAfterRestoreCapturesRestoredState | PASS |
| 14| TestRestoreUnknownSnapshotDoesNothing  | PASS   |
| 15| TestRestoreInvalidSnapshotIDs          | PASS   |
| 16| TestComplexStateTransitions            | PASS   |

## Final Test Output

All 16 tests passed on the first execution. The complete output is stored in:

test-output.txt

## Verification

- README modified: No
- SPEC modified: No
- Test files modified: No
- Dependency files modified: No
- Unexpected files created: No
- Accessed parent directories: No

## Self Verification

Before completing the benchmark, verify each item below.

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

- The SPEC.md was minimal ("Stack Machine Specification" with a note that the final spec and tests would be added before the first published run). The actual requirements were derived entirely from `solution_test.go`.
- All tests passed on the first attempt; no iteration was needed.
- Code was formatted with `gofmt -w .` as required by the README workflow.
