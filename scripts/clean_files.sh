#!/bin/bash

# Check if directory path and ignore file are provided
if [ $# -ne 2 ]; then
    echo "Usage: $0 <directory_path> <ignore_file>"
    exit 1
fi

DIR_PATH="$1"
IGNORE_FILE="$2"

# Check if directory exists
if [ ! -d "$DIR_PATH" ]; then
    echo "Error: Directory '$DIR_PATH' does not exist"
    exit 1
fi

# Check if ignore file exists
if [ ! -f "$IGNORE_FILE" ]; then
    echo "Error: Ignore file '$IGNORE_FILE' does not exist"
    exit 1
fi

# Function to convert wildcards in patterns to regex
convert_wildcards() {
    local pattern="$1"
    # Escape dots
    pattern=${pattern//./\\.}
    # Convert * to [^/]*
    pattern=${pattern//\*/[^/]*}
    # Convert ? to [^/]
    pattern=${pattern//\?/[ ^/]}
    echo "$pattern"
}

# Function to convert gitignore-style pattern to regex
convert_pattern() {
    local pattern="$1"
    # Skip empty lines and comments
    [[ -z "$pattern" || "$pattern" =~ ^[[:space:]]*# ]] && return
    # Check if it's a directory pattern (ends with /)
    local is_dir=false
    if [[ "$pattern" == */ ]]; then
        is_dir=true
    fi
    # Check if it's anchored (starts with /)
    local anchored=false
    if [[ "$pattern" == /* ]]; then
        anchored=true
        pattern=${pattern#/}  # Remove leading /
    fi
    # Convert wildcards
    local converted
    converted=$(convert_wildcards "$pattern")
    # Build regex based on pattern type
    local regex
    if $is_dir; then
        if $anchored; then
            regex="^$converted.*"
        else
            regex="(.*/)?$converted.*"
        fi
    else
        if $anchored; then
            regex="^$converted$"
        elif [[ "$pattern" != */* ]]; then
            regex="(.*/)?$converted$"
        else
            regex="^$converted$"
        fi
    fi
    echo "$regex"
}

# Read ignore patterns and convert them to regex before changing directories
declare -a IGNORE_REGEX
while IFS= read -r line || [[ -n "$line" ]]; do
    regex=$(convert_pattern "$line")
    [[ -n "$regex" ]] && IGNORE_REGEX+=("$regex")
done < "$IGNORE_FILE"

# Now change to the target directory
cd "$DIR_PATH" || exit 1

# Delete files that don't match any ignore pattern
find . -type f | while read -r file; do
    file=${file#./}  # Remove leading ./
    should_delete=true
    for pattern in "${IGNORE_REGEX[@]}"; do
        if [[ "$file" =~ $pattern ]]; then
            should_delete=false
            break
        fi
    done
    if $should_delete; then
        rm -f "$file"
    fi
done
