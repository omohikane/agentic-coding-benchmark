#!/usr/bin/env bash
set -uo pipefail

usage() {
  cat <<'EOF'
Usage:
  report-run.sh \
    --benchmark 001-versioned-store \
    --tool goose \
    --model qwen3.6-35b-heretic \
    [--prompt-file prompt.txt] \
    [--notes "First attempt"] \
    [--run-dir PATH] \
    [--results-dir PATH]

Example:
  ../../scripts/report-run.sh \
    --benchmark 001-versioned-store \
    --tool goose \
    --model qwen3.6-35b-heretic \
    --prompt-file prompt.txt
EOF
}

benchmark=""
tool=""
model=""
prompt_file=""
notes=""
run_dir="$PWD"
results_dir=""

while [[ $# -gt 0 ]]; do
  case "$1" in
  --benchmark)
    benchmark="${2:-}"
    shift 2
    ;;
  --tool)
    tool="${2:-}"
    shift 2
    ;;
  --model)
    model="${2:-}"
    shift 2
    ;;
  --prompt-file)
    prompt_file="${2:-}"
    shift 2
    ;;
  --notes)
    notes="${2:-}"
    shift 2
    ;;
  --run-dir)
    run_dir="${2:-}"
    shift 2
    ;;
  --results-dir)
    results_dir="${2:-}"
    shift 2
    ;;
  -h | --help)
    usage
    exit 0
    ;;
  *)
    echo "Unknown option: $1" >&2
    usage >&2
    exit 2
    ;;
  esac
done

if [[ -z "$benchmark" || -z "$tool" || -z "$model" ]]; then
  echo "Error: --benchmark, --tool, and --model are required." >&2
  usage >&2
  exit 2
fi

if [[ ! -d "$run_dir" ]]; then
  echo "Error: run directory does not exist: $run_dir" >&2
  exit 2
fi

run_dir="$(cd "$run_dir" && pwd)"

repo_root="$(
  git -C "$run_dir" rev-parse --show-toplevel 2>/dev/null ||
    true
)"

if [[ -z "$repo_root" ]]; then
  echo "Error: run directory is not inside a Git repository." >&2
  exit 2
fi

if [[ -z "$results_dir" ]]; then
  results_dir="$repo_root/results"
fi

mkdir -p "$results_dir"

slug="${benchmark}__${tool}__${model}"
report_file="$results_dir/${slug}.md"
test_output="$run_dir/test-output.txt"
solution_diff="$run_dir/solution.diff"
git_status_file="$run_dir/git-status.txt"

started_at="$(date --iso-8601=seconds)"
start_ns="$(date +%s%N)"

set +e
(
  cd "$run_dir"
  go test -v
) 2>&1 | tee "$test_output"
test_exit_code=${PIPESTATUS[0]}
set -e

end_ns="$(date +%s%N)"
finished_at="$(date --iso-8601=seconds)"
duration_ms=$(((end_ns - start_ns) / 1000000))

if [[ $test_exit_code -eq 0 ]]; then
  result="PASS"
else
  result="FAIL"
fi

git -C "$run_dir" status --short >"$git_status_file"

: >"$solution_diff"

if git -C "$run_dir" ls-files --error-unmatch solution.go >/dev/null 2>&1; then
  git -C "$run_dir" diff -- solution.go >"$solution_diff"
elif [[ -f "$run_dir/solution.go" ]]; then
  git diff --no-index /dev/null "$run_dir/solution.go" \
    >"$solution_diff" 2>/dev/null || true
fi

go_version="$(go version 2>/dev/null || echo unknown)"
git_commit="$(git -C "$repo_root" rev-parse --short HEAD 2>/dev/null || echo unknown)"
hostname_value="$(hostname 2>/dev/null || echo unknown)"
os_value="$(uname -srmo 2>/dev/null || uname -a)"
test_count="$(
  grep -c '^--- PASS:' "$test_output" 2>/dev/null || true
)"
fail_count="$(
  grep -c '^--- FAIL:' "$test_output" 2>/dev/null || true
)"

prompt_text="Not recorded."

if [[ -n "$prompt_file" ]]; then
  if [[ -f "$prompt_file" ]]; then
    prompt_text="$(cat "$prompt_file")"
  elif [[ -f "$run_dir/$prompt_file" ]]; then
    prompt_text="$(cat "$run_dir/$prompt_file")"
  else
    prompt_text="Prompt file not found: $prompt_file"
  fi
fi

changed_files="$(
  git -C "$run_dir" status --short |
    sed 's/^...//' |
    sed '/^[[:space:]]*$/d'
)"

if [[ -z "$changed_files" ]]; then
  changed_files="None"
fi

{
  cat <<EOF
# ${benchmark}: ${tool} / ${model}

## Summary

| Field | Value |
|---|---|
| Result | **${result}** |
| Benchmark | \`${benchmark}\` |
| Tool | \`${tool}\` |
| Model | \`${model}\` |
| Passed tests | ${test_count} |
| Failed tests | ${fail_count} |
| Test exit code | ${test_exit_code} |
| Duration | ${duration_ms} ms |
| Started | ${started_at} |
| Finished | ${finished_at} |

## Environment

| Field | Value |
|---|---|
| Host | \`${hostname_value}\` |
| OS | \`${os_value}\` |
| Go | \`${go_version}\` |
| Git revision | \`${git_commit}\` |
| Run directory | \`${run_dir#$repo_root/}\` |

## Prompt

\`\`\`text
${prompt_text}
\`\`\`

## Changed files

\`\`\`text
${changed_files}
\`\`\`

## Git status

\`\`\`text
$(cat "$git_status_file")
\`\`\`

## Test output

\`\`\`text
$(cat "$test_output")
\`\`\`

## Solution diff

\`\`\`diff
$(cat "$solution_diff")
\`\`\`

## Notes

${notes:-No additional notes.}
EOF
} >"$report_file"

echo
echo "Result: $result"
echo "Report: $report_file"
echo "Test output: $test_output"
echo "Solution diff: $solution_diff"

exit "$test_exit_code"
