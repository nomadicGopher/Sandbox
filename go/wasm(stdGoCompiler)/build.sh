#!/bin/bash
LOG_FILE="build.log"
touch $LOG_FILE 2>&1 | tee /dev/stderr

echo "----------------------------------------------------------------------------------------------------" | tee -a $LOG_FILE
echo "Build process started at: $(date)" | tee -a $LOG_FILE

echo "Removing web/go.wasm if it exists..." | tee -a $LOG_FILE
rm -f web/go.wasm 2>&1 | tee -a $LOG_FILE >&2

echo "Building a new web/go.wasm..." | tee -a $LOG_FILE
GOOS=js GOARCH=wasm go build -o=web/go.wasm -buildvcs=false main.go 2>&1 | tee -a $LOG_FILE >&2
status=$?
if [ $status -ne 0 ]; then
    echo "Failed to build web/go.wasm" | tee -a $LOG_FILE >&2
    exit 1
fi
echo "web/go.wasm was built." | tee -a $LOG_FILE

if [ ! -f web/wasm_exec.js ]; then
    echo "Fetching a new web/wasm_exec.js..." | tee -a $LOG_FILE 
    cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" web/ 2>&1 | tee -a $LOG_FILE >&2  
    echo "web/wasm_exec.js was fetched from \$GOROOT/misc/wasm/" | tee -a $LOG_FILE
else
    echo "web/wasm_exec.js already exists. Skipping fetch." | tee -a $LOG_FILE
fi

echo "Build process ended at: $(date)" | tee -a $LOG_FILE
exit
