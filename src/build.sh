#!/bin/bash

go build -o mac-launcher-rename

env GOOS=windows GOARCH=amd64 go build -o windows-launcher-rename.exe

mv *-launcher-rename* ../run