# AI Coding Agent Benchmark

This directory contains a single benchmark task for evaluating AI coding agents.

The benchmark consists of a specification, one or more tests, and a project skeleton. Your objective is to implement the required functionality so that all tests pass.

This benchmark evaluates implementation ability only. Scoring is performed separately by a dedicated Judge Agent.

When in doubt, prefer reporting uncertainty over making assumptions.

The benchmark values factual execution more than optimistic reporting.

---

# Objective

Complete the benchmark by performing the following steps:

1. Understand the specification.
2. Implement the required functionality.
3. Execute the project's test suite.
4. Fix any failing tests.
5. Verify the final implementation.
6. Produce execution logs and a run report.

---

# Files

A benchmark directory typically contains:

- `README.md`
  - General benchmark instructions (this file)

- `SPEC.md`
  - Task-specific requirements

- Test files
- Project configuration files
- Skeleton source files

The exact filenames depend on the programming language.

Examples:

- Go
  - `solution.go`
  - `solution_test.go`
  - `go.mod`

- Python
  - `solution.py`
  - `test_solution.py`
  - `pyproject.toml`

- Rust
  - `src/lib.rs`
  - `tests/`
  - `Cargo.toml`

---

# Working Directory

Only work inside the current benchmark directory.

Do **not** access:

- parent directories
- other benchmark directories
- other run directories
- scoring tools
- previous benchmark results
- reference implementations
- solution repositories

Assume that anything outside the current directory is unavailable.

---

# Protected Files

Unless the specification explicitly allows it, do **not** modify:

- `README.md`
- `SPEC.md`
- test files
- project configuration files
- provided input data

Examples of protected test files include:

- `*_test.go`
- `test_*.py`
- `*_test.py`
- `*.spec.js`
- `*.test.js`
- `*.spec.ts`
- `*.test.ts`

Examples of protected project files include:

- `go.mod`
- `go.sum`
- `Cargo.toml`
- `Cargo.lock`
- `package.json`
- `package-lock.json`
- `pyproject.toml`
- `requirements.txt`

Only modify files that are necessary to implement the solution.

---

# Prohibited Actions

The following actions are not allowed:

- modifying tests
- disabling or skipping tests
- changing the specification
- bypassing tests with hard-coded outputs
- faking test execution
- reporting results that were not actually produced
- reading data outside the benchmark directory
- adding unnecessary dependencies
- generating your own score

A dedicated Judge Agent performs scoring later.

Your responsibility is only to implement the solution and produce an accurate execution report.

---

# Expected Workflow

## 1. Inspect

Read:

- `SPEC.md`
- all test files

Understand:

- required behavior
- project structure
- implementation target

---

## 2. Implement

Implement the required functionality.

Guidelines:

- make the smallest reasonable changes
- preserve the existing project structure
- avoid unrelated refactoring
- follow existing coding style

---

## 3. Format

If a standard formatter is available for the language, run it.

Examples:

Go

```bash
gofmt -w .
```

Python

```bash
ruff format .
```

Rust

```bash
cargo fmt
```

JavaScript / TypeScript

```bash
npx prettier --write .
```

---

## 4. Test

Execute the project's test suite.

Examples:

Go

```bash
go test -v ./...
```

Python

```bash
pytest -v
```

Rust

```bash
cargo test
```

JavaScript

```bash
npm test
```

If tests fail:

1. inspect the failure
2. modify the implementation
3. run the tests again

Repeat until no further progress can be made.

---

## 5. Verify

Before finishing, verify that:

- the implementation compiles (if applicable)
- formatting has been applied
- tests have been executed
- protected files were not modified
- no unnecessary files were created

---

## 6. Record Results

Save the final test output to:

```text
test-output.txt
```

For example:

```bash
go test -v ./... 2>&1 | tee test-output.txt
```

Use the equivalent command for the project's language if necessary.

Then generate:

```text
run-report.md
```

---

# Required Outputs

At the end of the benchmark, the directory should contain:

- implementation files
- `test-output.txt`
- `run-report.md`

Do not create additional files unless they are required by the implementation.

---

# Run Report Format

Create `run-report.md` using the following structure.

```markdown
# Run Report

## Summary

- Result: PASS / FAIL
- Language:
- Implementation Files:
- Final Test Command:
- Final Test Status:
- Implementation Attempts:
  Record the actual number of implementation iterations.
  If this was not tracked during execution,
  write:
  Not Recorded
- Test Executions:

## Commands Executed

List every command that was executed during the benchmark.

Example:

1. ls

2. cat README.md

3. cat SPEC.md

4. go test -v ./...

5. gofmt -w .

6. go test -v ./...

## Files Changed

List every file that was created or modified.

## Implementation Summary

Briefly describe the implementation.

## Test Progression

### Test Run 1

- Command:
- Result:
- Main Failures:
- Changes Made:

Repeat for each test execution if applicable.

## Final Test Output

Summarize the final test result.

The complete output is stored in:

test-output.txt

## Verification

- README modified: Yes / No
- SPEC modified: Yes / No
- Test files modified: Yes / No
- Dependency files modified: Yes / No
- Unexpected files created: Yes / No
- Accessed parent directories: Yes / No

## Self Verification

Before completing the benchmark, verify each item below.

| Item                                                       | Status   |
| ---------------------------------------------------------- | -------- |
| Read and understood `SPEC.md`                              | Yes / No |
| Implemented the required functionality                     | Yes / No |
| Executed the final test suite                              | Yes / No |
| Saved the final test output to `test-output.txt`           | Yes / No |
| Generated `run-report.md`                                  | Yes / No |
| Modified any protected files                               | Yes / No |
| Accessed parent directories or other benchmark directories | Yes / No |
| Created unnecessary files                                  | Yes / No |

Do not guess.

If you cannot verify an item from actual execution, mark it as **Unknown** and explain why in the **Notes** section.

This checklist is a factual verification of the completed work, not a summary or an estimate.

## Notes

Record any assumptions, uncertainties, or remaining concerns.
```

---

## Reporting Rules

The run report is a factual execution log.

Everything written in the report must be based on actions that were actually performed.

Never estimate or infer information that was not explicitly recorded during execution.

If a value is unknown, write:

- Not Recorded

instead of guessing.

Examples include:

- execution time
- implementation attempts
- number of test executions
- command history
- token usage
- benchmark score

Accuracy is more important than completeness.

When generating `run-report.md`, do not summarize from memory.

Instead, use the files generated during execution and the commands that were actually executed.

The report should be reproducible from the benchmark artifacts.

Treat `run-report.md` as an execution log, not as a narrative summary.

Prefer omitting information over guessing.

A missing fact is acceptable.
An invented fact is not.

---

# Completion Criteria

The benchmark is complete when:

- the implementation has been produced
- the final test suite has been executed
- the final test output has been saved to `test-output.txt`
- `run-report.md` has been generated
- protected files remain unchanged
- all work was performed within the current directory

If the implementation is incomplete or tests still fail, generate the report anyway and record the final status honestly as **FAIL**.
