#!/bin/bash

# Create a test PE file for Mali analyzer
# This script tries multiple methods to create a valid PE file

cd "$(dirname "$0")/.." || exit 1

echo "Creating test PE file..."

# Method 1: Try using Go to create minimal PE
if command -v go &> /dev/null; then
    echo "Method 1: Creating minimal PE with Go..."
    go run scripts/create_minimal_pe.go 2>/dev/null
    if [ -f samples/minimal_pe.exe ]; then
        if head -c 2 samples/minimal_pe.exe | grep -q "MZ"; then
            echo "✓ Created minimal_pe.exe (PE format)"
            cp samples/minimal_pe.exe samples/test_pe.exe
            rm -f samples/minimal_pe.exe
            file samples/test_pe.exe
            exit 0
        fi
    fi
fi

cat > samples/test_pe.c << 'EOF'
#include <stdio.h>
int main() {
    printf("Test PE file for Mali analyzer\n");
    return 0;
}
EOF

if command -v x86_64-w64-mingw32-gcc &> /dev/null; then
    echo "Compiling with mingw-w64 (x86_64)..."
    cd samples
    x86_64-w64-mingw32-gcc -o test_pe.exe test_pe.c 2>&1
    cd ..
    if [ $? -eq 0 ]; then
        echo "✓ Created test_pe.exe (PE format)"
        file test_pe.exe
    else
        echo "✗ Failed to compile with mingw-w64"
        exit 1
    fi
elif command -v i686-w64-mingw32-gcc &> /dev/null; then
    echo "Compiling with mingw-w64 (i686)..."
    cd samples
    i686-w64-mingw32-gcc -o test_pe.exe test_pe.c 2>&1
    cd ..
    if [ $? -eq 0 ]; then
        echo "✓ Created test_pe.exe (PE format)"
        file test_pe.exe
    else
        echo "✗ Failed to compile with mingw-w64"
        exit 1
    fi
else
    echo "✗ mingw-w64 not found. Installing..."
    if command -v dnf &> /dev/null; then
        echo "Installing mingw64-gcc (requires sudo)..."
        sudo dnf install -y mingw64-gcc 2>&1 | tail -5
    elif command -v apt-get &> /dev/null; then
        echo "Installing mingw-w64 (requires sudo)..."
        sudo apt-get install -y mingw-w64 2>&1 | tail -5
    else
        echo "Please install mingw-w64 manually:"
        echo "  Fedora: sudo dnf install mingw64-gcc"
        echo "  Ubuntu/Debian: sudo apt install mingw-w64"
        exit 1
    fi
    
    if command -v x86_64-w64-mingw32-gcc &> /dev/null; then
        echo "Compiling with newly installed mingw-w64..."
        cd samples
        x86_64-w64-mingw32-gcc -o test_pe.exe test_pe.c 2>&1
        cd ..
        if [ $? -eq 0 ]; then
            echo "✓ Created test_pe.exe (PE format)"
            file test_pe.exe
        else
            echo "✗ Failed to compile"
            exit 1
        fi
    else
        echo "✗ Installation failed or compiler not in PATH"
        exit 1
    fi
fi

