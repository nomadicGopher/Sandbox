#!/bin/bash

# Define the track you want to download exercises for
TRACK="go"

# Define the directory where you want to save the exercises
DOWNLOAD_DIR="$HOME/exercism/$TRACK"

echo "Attempting to download all active '$TRACK' exercises to: $DOWNLOAD_DIR"
echo "Please ensure you have the exercism CLI installed and configured with your API token."
echo ""

# Create the download directory if it doesn't exist
mkdir -p "$DOWNLOAD_DIR"

# Change to the download directory
cd "$DOWNLOAD_DIR" || { echo "Error: Could not change to directory $DOWNLOAD_DIR"; exit 1; }

# Get a list of all active exercises for the specified track
# The `exercism list` command is used, and we filter for the desired track
# We then extract the exercise slug (e.g., "hello-world", "two-fer")
# This assumes that `exercism list` output format is consistent.
echo "Fetching list of active '$TRACK' exercises..."
EXERCISES=$(exercism list | grep "track:$TRACK" | awk '{print $NF}' | tr -d '()')

if [ -z "$EXERCISES" ]; then
    echo "No active '$TRACK' exercises found or there was an issue fetching the list."
    echo "Please check your exercism CLI configuration and internet connection."
    exit 0
fi

echo "Found the following active exercises for '$TRACK':"
echo "$EXERCISES"
echo ""

# Loop through each exercise and download it
for EXERCISE in $EXERCISES; do
    echo "Downloading '$EXERCISE'..."
    exercism download --track="$TRACK" --exercise="$EXERCISE"
    if [ $? -ne 0 ]; then
        echo "Warning: Failed to download '$EXERCISE'. Moving to the next exercise."
    fi
    echo "-------------------------------------"
done

echo "All active '$TRACK' exercises have been processed."
echo "You can find them in: $DOWNLOAD_DIR"
