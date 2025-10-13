#!/bin/bash

# Exit on error
set -e

# Use python3 if available, fallback to python
PYTHON_BIN=$(command -v python3 || command -v python)
if [ -z "$PYTHON_BIN" ]; then
    echo "Python is not installed."
    exit 1
fi

# Create a virtual environment with access to system site packages
$PYTHON_BIN -m venv venv --system-site-packages

# Activate the virtual environment
if [ -f "venv/bin/activate" ]; then
    # Linux/macOS
    source venv/bin/activate
elif [ -f "venv/Scripts/activate" ]; then
    # Windows (Git Bash, WSL, etc.)
    source venv/Scripts/activate
else
    echo "Could not find the virtual environment activation script."
    exit 1
fi

# Upgrade pip
pip install --upgrade pip

# Install the required packages
pip install pytest-cache pytest-subtests pylint
