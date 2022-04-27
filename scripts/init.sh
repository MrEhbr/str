#!/usr/bin/env bash

root=$(git rev-parse --show-toplevel)
project=$(basename $root)
for file in $(git grep golang-repo-template | cut -d: -f1 | sort -u); do
    sed -i '' s/golang-repo-template/$project/g $file
done

mv $root/cmd/golang-repo-template $root/cmd/$project

# delete script it not need anymore
rm $0
