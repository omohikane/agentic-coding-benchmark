# Agentic Coding Benchmark

A reproducible benchmark for evaluating AI coding models and coding agents on practical repository tasks.

The benchmark focuses on more than code generation. Each run evaluates whether an AI system can:

- inspect a repository
- read task instructions and specifications
- create or modify files
- run an existing test suite
- interpret failures
- repair its implementation within a limited number of attempts
- preserve files that it was instructed not to change

## Goals

This repository is designed to compare both:

1. **Models** such as local and cloud LLMs
2. **Tools** such as Goose, OpenCode, Claude Code, Codex CLI, Aider, and Gemini CLI

The benchmark avoids relying only on well-known algorithm exercises. Tasks should resemble small but realistic software engineering assignments with explicit specifications and hidden edge cases.

## Repository layout

```text
benchmarks/   Benchmark tasks
results/      Published run reports
scripts/      Helper scripts for repeatable runs
docs/         Benchmark rules and contribution notes
```

## Running a benchmark

Each benchmark directory contains its own `README.md` with the exact instruction that should be given to the coding agent.

A typical run is:

1. Copy the benchmark directory into a clean working directory.
2. Start the selected coding tool in that directory.
3. Give it only the instruction written in the benchmark README.
4. Record the model, tool, configuration, elapsed time, test attempts, edits, and final result.
5. Save the report under `results/`.

## Fairness rules

- Use a fresh working directory for every run.
- Do not reuse files created by another model or tool.
- Do not manually repair the implementation during a run.
- Do not change the provided tests.
- Record retries, failures, and human interventions.
- Record model quantization and runtime settings for local models.
- Record tool version and permission settings.

See [`docs/methodology.md`](docs/methodology.md) for the full protocol.

## Status

This project is experimental. Benchmark specifications and scoring may evolve while the initial task set is being developed.
