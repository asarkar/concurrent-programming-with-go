#!/bin/bash

set -e

no_test=0
no_lint=0

while (( $# > 0 )); do
   case "$1" in
   	--help)
			printf "run.sh [OPTION]... [DIR]\n"
			printf "options:\n"
			printf "\t--help			Show help\n"
			printf "\t--no-test		Skip tests\n"
			printf "\t--no-lint		Skip linting\n"
			exit 0
      	;;
      --no-test)
			no_test=1
			shift
      	;;
      --no-lint)
			no_lint=1
			shift
			;;
		*)
			break
	      ;;
   esac
done

if (( no_test == 0 )); then
  if [[ -z "$1" ]]; then
    go test -v ./...
  else
    go test -v ./"$1"/...
  fi
fi

if (( no_lint == 0 )); then
	if [[ -z "$1" ]]; then
    go vet ./...
  else
    go vet ./"$1"/...
  fi
fi