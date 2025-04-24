#!/bin/bash

# Constants
LOG_FILE="build.log"
WASM_FILE="web/go.wasm"
WASM_EXEC_JS="web/wasm_exec.js"

# Define a logging function
log() {
    local level="$1"
    local message="$2"
    echo "[$level] $message" | tee -a $LOG_FILE
}

# Initialize log file
touch $LOG_FILE 2>&1 | tee /dev/stderr

# Log start of build process
echo "====================================================================================================" | tee -a $LOG_FILE
echo "INFO" "Build process started at: $(date)" | tee -a $LOG_FILE
echo "====================================================================================================" | tee -a $LOG_FILE

# Remove existing wasm file
log "STEP" "Removing $WASM_FILE if it exists..."
rm -f $WASM_FILE 2>&1 | tee -a $LOG_FILE >&2

# Build new wasm file
log "STEP" "Building a new $WASM_FILE..."
GOOS=js GOARCH=wasm go build -o=$WASM_FILE -buildvcs=false main.go 2>&1 | tee -a $LOG_FILE >&2
if [ $? -ne 0 ]; then
    log "ERROR" "Failed to build $WASM_FILE"
    exit 1
fi
log "SUCCESS" "$WASM_FILE was built."

# Determine OS-specific path separator
if [[ "$OSTYPE" == "msys" || "$OSTYPE" == "cygwin" || "$OSTYPE" == "win32" ]]; then
    PATH_SEPARATOR='\\'
else
    PATH_SEPARATOR='/'
fi

# Fetch web/wasm_exec.js if not present
log "STEP" "Checking for $WASM_EXEC_JS..."
if [ ! -f $WASM_EXEC_JS ]; then
    log "INFO" "$WASM_EXEC_JS not found. Fetching from Go resources..."
    
    # Locate wasm_exec.js
    GOROOT=$(go env GOROOT)
    WASM_EXEC_PATH="$GOROOT${PATH_SEPARATOR}misc${PATH_SEPARATOR}wasm${PATH_SEPARATOR}wasm_exec.js"

    # Normalize path for display
    if [[ "$OSTYPE" == "msys" || "$OSTYPE" == "cygwin" || "$OSTYPE" == "win32" ]]; then
        DISPLAY_WASM_EXEC_PATH=$(echo "$WASM_EXEC_PATH" | sed 's/\\\\/\\/g')
    else
        DISPLAY_WASM_EXEC_PATH="$WASM_EXEC_PATH"
    fi
    
    # Check if go ... wasm_exec.js exists
    if [ ! -f "$WASM_EXEC_PATH" ]; then
        log "ERROR" "wasm_exec.js not found at expected location: $DISPLAY_WASM_EXEC_PATH"
        return 1
    fi
    
    cp "$WASM_EXEC_PATH" web/ 2>&1 | tee -a $LOG_FILE >&2
    log "SUCCESS" "$WASM_EXEC_JS was fetched from $DISPLAY_WASM_EXEC_PATH"
else
    log "INFO" "$WASM_EXEC_JS already exists. Skipping fetch."
fi

# Log end of build process
echo "====================================================================================================" | tee -a $LOG_FILE
echo "Build process ended at: $(date)" | tee -a $LOG_FILE
echo "====================================================================================================" | tee -a $LOG_FILE
exit
