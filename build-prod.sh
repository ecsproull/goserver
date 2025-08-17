#!/bin/bash

echo "🔨 Building Go server for production..."

# Clean previous build
if [ -f "goserver_prod" ]; then
    echo "Removing previous build..."
    rm goserver_prod
fi

# Build with production settings
echo "Compiling binary..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o goserver_prod

# Check if build was successful
if [ $? -eq 0 ]; then
    echo "✅ Build successful!"
    echo "📊 Binary information:"
    ls -lh goserver_prod
    echo ""
    echo "🚀 Ready for deployment:"
    echo "   scp goserver_prod user@server:/path/to/destination/"
else
    echo "❌ Build failed!"
    exit 1
fi
