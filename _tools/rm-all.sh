#!/bin/bash

for file in *; do
    # _toolsは削除しない
    if [[ $file == _* ]]; then
        continue
    fi

    if [[ -d $file ]]; then
        rm -rf "$file"
    fi
done
