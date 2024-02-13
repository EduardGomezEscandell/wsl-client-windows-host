#!/bin/bash
set -eu

# Build the server for Windows
cd server
GOOS="windows" go build
cd - > /dev/null

# Build the client for Linux
cd client
GOOS="linux" go build
cd - > /dev/null