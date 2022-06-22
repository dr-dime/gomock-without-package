#!/bin/bash

mockgen_all() {
  for pkg in p1 p2; do
    find "${pkg}" -name *.go -and -not -name *_test.go -print -exec \
      "${BASH_SOURCE}" {} \;
  done
}

mockgen_file() {
  local src="$1"
  local dest="${src%.go}_mockgen_test.go"
  mockgen -source="${src}" -destination="${dest}" \;
}

main() {
  if (( $# > 0 )); then
    mockgen_file "$@"
  else
    mockgen_all
  fi
}

main "$@"
