#!/bin/sh
TARGETS="$1"

if [ -z "$TARGETS" ]; then
    echo "Usage checker.sh startdir"
    exit 1
fi

for i in $TARGETS; do
    find $i -type f -print0 | xargs -0 md5sum
done

