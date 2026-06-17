#!/bin/bash

echo "=== Running no singleton ==="

go run no_singleton/main.go

echo ""

echo "Running basic singleton ==="

go run basic/main.go
echo ""

echo "Running singleton with sync.Once ==="

go run mutex/main.go