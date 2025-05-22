#!/bin/bash
rm *.log
rm *transformed*

go run . # Relative path using default flag
#go run . -inPath="$GOPATH\src\Sandbox\go\data-sources\InputFile\X12\sample.278" # Windows path
#go run . -inPath="$GOPATH/src/Sandbox/go/data-sources/InputFile/X12/sample.278" # Linux path