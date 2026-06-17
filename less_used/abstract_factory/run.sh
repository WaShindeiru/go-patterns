#!/bin/bash

echo "=== Running Java ==="

javac Main.java
java Main

rm -f *.class

echo ""
echo "=== Running Go ==="

go run main.go