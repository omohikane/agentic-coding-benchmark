#!/bin/sh
set -eu

if [ "$#" -ne 4 ]; then
    echo "usage: $0 <benchmark-id> <tool> <model> <date>" >&2
    exit 1
fi

benchmark_id=$1
tool=$2
model=$3
date=$4

safe_tool=$(printf '%s' "$tool" | tr ' /' '--')
safe_model=$(printf '%s' "$model" | tr ' /' '--')
output="results/${benchmark_id}__${safe_tool}__${safe_model}__${date}.md"

if [ -e "$output" ]; then
    echo "file already exists: $output" >&2
    exit 1
fi

cat > "$output" <<EOF2
# Run Result

## Summary

- Benchmark: $benchmark_id
- Benchmark version:
- Date: $date
- Final status:

## Model

- Model: $model
- Provider or runtime:
- Quantization:
- Context size:
- Relevant inference settings:

## Coding tool

- Tool: $tool
- Version:
- Agent or mode:
- Permission configuration:

## Environment

- Operating system:
- CPU:
- GPU:
- RAM:

## Run metrics

- Elapsed time:
- Test executions:
- Implementation revisions:
- Human interventions:

## Test progression

1. Initial test result:
2. Second test result:
3. Final test result:

## Observed behavior


## Final notes

EOF2

printf '%s\n' "$output"
