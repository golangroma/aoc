#!/bin/bash

set -euo pipefail

source .env

YEAR=$1
DAY=$(printf %02d $2)

mkdir -p ./${YEAR}/day${DAY}/${USERNAME}
cp ./cli/template/* ./${YEAR}/day${DAY}/${USERNAME}
cd ./${YEAR}/day${DAY}/${USERNAME}
go mod init github.com/golangroma/aoc/${YEAR}/day${DAY}/${USERNAME}
echo "replace github.com/golangroma/aoc/cli => ../../../cli" >> go.mod
go mod tidy

if [[ -z ${session} ]]; then
    echo "Missing 'Advent of Code' session. Cannot download input."
else
    curl -s --cookie "session=${session}" https://adventofcode.com/$1/day/$(printf %d $2)/input > input.txt
fi

echo "You can now type 'cd ./${YEAR}/day${DAY}/${USERNAME}' and work on your solution!"
echo "To execute it, type 'go run .'"
