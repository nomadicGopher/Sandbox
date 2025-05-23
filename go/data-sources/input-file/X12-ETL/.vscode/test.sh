#!/bin/bash

# Remove log & data test files
rm -f ./*.log
rm -f ./*transformed*

cd ..

go run . # Relative path using default flag
#go run . -inPath="$GOPATH\src\Sandbox\go\data-sources\input-file\X12-ETL\.vscode\hhpcr_x12n278.txt" # Windows path
#go run . -inPath="$GOPATH/src/Sandbox/go/data-sources/input-file/X12-ETL/.vscode/hhpcr_x12n278.txt" # Linux path