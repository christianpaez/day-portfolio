#!/bin/bash
[ $# -ne 2 ] && { echo "Usage: $0 <zipfile> <imagename>"; exit 1; }

ZIP="$1"
NAME="$2"
TEMP="/tmp/unzip_$$"
mkdir -p "$TEMP"
unzip -q "$ZIP" -d "$TEMP"
find "$TEMP" -name "*.png" -o -name "*.jpg" -o -name "*.jpeg" | head -1 | while read -r IMG; do
    mv "$IMG" "/home/christianpaez/Documents/day-portfolio/assets/img/${NAME}.${IMG##*.}"
done
rm -rf "$TEMP"
