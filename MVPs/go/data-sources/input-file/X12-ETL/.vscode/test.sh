#!/bin/bash

# Remove log & data test files
rm -f ./*.log
rm -f ./*transformed*

cd ..

go run . -inPath=".vscode/hhpcr_x12n278.txt" -log=true