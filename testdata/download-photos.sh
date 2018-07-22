#!/usr/bin/env bash

set -e
set -x

main() {
    local width=1600
    local height=1200
    local count=0
    while [[ "${count}" -lt 50 ]]; do
        curl "http://placekitten.com/${width}/${height}" > "testdata/benchmark/${count}.jpg"
        count=$(expr $count + 1)
        width=$(expr $width + 1)
        height=$(expr $height + 1)
    done
}

main "$@"
