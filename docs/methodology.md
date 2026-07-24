# Benchmark Methodology

## What is being measured

The benchmark separates three related abilities:

1. **Coding ability**
   - correctness of the initial implementation
   - specification comprehension
   - code quality

2. **Agentic execution ability**
   - repository inspection
   - file creation and editing
   - test execution
   - failure diagnosis
   - repair behavior

3. **Tool integration quality**
   - reliability of file operations
   - command execution
   - context handling
   - permission behavior

## Required run metadata

Every result should include:

- benchmark ID and version
- date
- model name
- model provider or runtime
- quantization, context size, and inference settings when applicable
- coding tool and version
- system prompt or agent mode
- permission configuration
- operating system and hardware
- elapsed time
- number of test executions
- number of implementation revisions
- final test result
- whether human intervention occurred

## Standard run protocol

1. Create a clean copy of the benchmark directory.
2. Do not expose earlier model outputs to the system under test.
3. Start the coding tool inside the benchmark directory.
4. Give the tool only the task instruction defined by the benchmark.
5. Allow the tool to inspect repository files.
6. Allow the tool to create the requested implementation file.
7. Allow it to run the supplied tests.
8. Stop the run when one of the following occurs:
   - all tests pass
   - the tool declares completion
   - the configured attempt limit is reached
   - the tool becomes stuck or repeats without useful progress
9. Save the complete result report.

## Suggested limits

Initial benchmark runs should use:

- maximum test executions: 4
- maximum implementation revisions: 3
- external network access: disabled
- dependency policy: standard library only unless the task says otherwise

## Scoring

A suggested 100-point score:

| Category | Points |
|---|---:|
| Initial implementation correctness | 35 |
| Final correctness | 25 |
| Specification compliance | 15 |
| Repair efficiency | 10 |
| Code quality | 10 |
| Instruction compliance | 5 |

The raw evidence should always be published alongside the score.
