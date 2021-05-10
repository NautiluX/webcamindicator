#!/bin/bash

set -euo pipefail

if [ -z "$GOPATH" ]; then
    echo GOPATH environment variable not set
    exit
fi

if [ ! -e "$GOPATH/bin/2goarray" ]; then
    echo "Installing 2goarray..."
    go get github.com/cratonica/2goarray
fi

INPUT=$1
OUTPUT=$2
VARNAME=$3
echo Generating "$OUTPUT"
echo "//+build linux darwin" > "$OUTPUT"
echo >> "$OUTPUT"
"$GOPATH/bin/2goarray" "$VARNAME" icon >> "$OUTPUT" < "$INPUT"
echo Finished
