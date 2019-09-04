#/usr/bin/env bash

# /proc/sys/fs/inotify/max_user_watches default: 8192

server() {
  local go_exec_path="${1-}"
  local dir="$(dirname "$BASH_SOURCE")/../"
  local this_dir="$(cd "$dir" && pwd)"

  local env_source_fpath="$this_dir/env.source.sh"

  . "$env_source_fpath" || exit

  local go_repo_build_src_dir="$GOPATH/src/bitbucket.org/server-monitor/checkrlight"
  local watch_dir='.'
  local exclude_from_inotify='app\/javascript\/app|\.git\/|main\/filters|test|TODO\.notes'
  local js_reloader_file="$this_dir/app/javascript/app/src/devServerReload.js"

  local default_go_exec_path='main/http/'

  [ -n "$go_exec_path" ] || go_exec_path="$default_go_exec_path"

  local go_repo_build_src_fpath="$go_repo_build_src_dir/$go_exec_path"
  local output_bin="bin/$(basename --suffix .go "$go_exec_path")"
  local output_bin_fpath="$this_dir/$output_bin"
  local pgrep_key="$(basename "$output_bin_fpath")"

  validate_files go_repo_build_src_fpath js_reloader_file

  while true; do
    if ! go build -o "$output_bin_fpath" "$go_repo_build_src_fpath"; then
      local err="$?"
      generate_js_reloader "$js_reloader_file"
      exit "$?"
    fi

    "$output_bin_fpath" &

    # Sleep for a little bit to let the back end get ready to accept connections before
    #   forcing reload on the front end.
    sleep 4

    generate_js_reloader "$js_reloader_file"

    inotifywait --recursive --event modify \
      --exclude "$exclude_from_inotify" \
      "$watch_dir"
  done
}

validate_files() {
  local vname=''
  local fname=''
  for vname in "$@"; do
    fname="${!vname}"
    if ! [ -s "$fname" ]; then
      printf '%s\n%s\n' \
        "... ERROR: file doesn't exist or is empty:" \
        "  '$vname' => '$fname'" >&2
      exit 1
    fi
  done
}

generate_js_reloader() {
  local js_reloader_file="${1:?ERROR, must pass JS reloader file.}"

  echo "export default () => { console.log('$(date) ... dev server reload.'); };" > "$js_reloader_file"
}

set -o errexit
set -o pipefail
set -o nounset

server "$@"
