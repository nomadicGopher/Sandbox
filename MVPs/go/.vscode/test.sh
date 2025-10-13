#!/bin/bash

# Remove log & data test files
rm -f ./*.log
rm -f ./{datafile}

cd ..

go run .