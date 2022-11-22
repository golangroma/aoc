#!/bin/bash

set -euo pipefail

source .env

YEAR=$1
DAY=$(printf %02d $2)

mkdir -p ./$YEAR/day$DAY/$USERNAME
cp ./template/* ./$YEAR/day$DAY/$USERNAME
cd ./$YEAR/day$DAY/$USERNAME

if [[ -z $session ]]; then
    echo "Missing 'Advent of Code' session. Cannot download input."
else
    curl -s --cookie "session=$session" https://adventofcode.com/$1/day/$(printf %d $2)/input > input.txt
fi

echo "You can now run 'cd ./$YEAR/day$DAY/$USERNAME' and work on your solution!"
