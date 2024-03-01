#!/bin/bash

# This is the alias that will be created in the sistem, you can change it, if you want
OUTPUT_FILE="wlopt"

# Building the program
echo "Compiling the main program..."
go build -o "$OUTPUT_FILE" main.go

# Updating system paths
    echo "Updating system paths ..."
    echo "alias wlopt=\"$PWD/wlopt\"" >> ~/.bashrc
    echo "alias wlopt=\"$PWD/wlopt\"" >> ~/.zshrc

echo "Done!"
echo "Please now run the following commands to reload your terminal or open a new one "
echo "source ~/.bashrc "
echo "source ~/.zshrc"
