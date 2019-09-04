#!/usr/bin/env bash

kill-go-server() {
  # The go bin base name.
  local pgrep_key='http'

  local pid=''
  declare -a pids=()
  while read pid; do
    pids+=( "$pid" )
  done < <(pgrep "$pgrep_key")

  local count="${#pids[@]}"
  if [ "$count" -ne 1 ]; then
    echo "ERROR: '$pgrep_key' procs count '$count' should be == 1, kill manually..." >&2
    ps aux | grep "$pgrep_key"
    exit 1
  fi

  pid="${pids[@]}"
  echo "... killing '$pgrep_key' '$pid'..."
  ps -p "$pid"
  kill "$pid"
  ps -p "$pid" >/dev/null && echo "ERROR: unable to kill '$pgrep_key'/'$pid', kill manually..." >&2
  pgrep 'inotifywait' && echo 'ERROR: inotifywait proc present, kill manually...' >&2
}

set -o errexit
set -o pipefail
set -o nounset

kill-go-server "$@"
