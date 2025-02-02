# !/bin/bash

KOBO_PATH="/Volumes/KOBOeReader"

if [ ! -d "$KOBO_PATH" ]; then
    echo "Error: Kobo eReader not found at $KOBO_PATH"
    echo "Please ensure your Kobo is connected and mounted"
    exit 1
fi

rm -rf "$KOBO_PATH"/*

cp -r ./kobo/.??* "$KOBO_PATH"/

diskutil eject "$KOBO_PATH"

echo "Kobo eReader reset setup complete"