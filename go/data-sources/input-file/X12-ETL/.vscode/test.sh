#!/bin/bash

# Remove log & data test files
rm -f ./*.log
rm -f ./*transformed*

cd ..

go run .