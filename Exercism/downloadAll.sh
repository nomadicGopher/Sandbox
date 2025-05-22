#!/bin/bash

TRACK_SLUG="go"
EXERCISES_FILE="exercises.txt"

echo "Fetching list of Go exercises..."

# This command scrapes the exercise slugs from the Exercism website
# It might break if Exercism changes their website structure.
# A more robust solution might involve using the Exercism API if available.
curl -LSs "https://exercism.org/tracks/${TRACK_SLUG}/exercises" | \
  grep "/tracks/${TRACK_SLUG}/exercises/" | \
  awk '{print $3}' | \
  cut -d/ -f5 | \
  cut -d\" -f1 > "$EXERCISES_FILE"

if [ ! -s "$EXERCISES_FILE" ]; then
  echo "Failed to retrieve exercise list. Please check the curl command and Exercism website structure."
  exit 1
fi

echo "Downloading exercises for the '$TRACK_SLUG' track..."
while IFS= read -r exercise_slug; do
  if [ -n "$exercise_slug" ]; then
    echo "Downloading: $exercise_slug"
    exercism download --exercise="$exercise_slug" --track="$TRACK_SLUG"
    if [ $? -ne 0 ]; then
      echo "Warning: Failed to download $exercise_slug. It might be locked or an error occurred."
    fi
  fi
done < "$EXERCISES_FILE"

echo "Finished attempting to download all exercises."

rm "$EXERCISES_FILE" # Clean up the temporary file