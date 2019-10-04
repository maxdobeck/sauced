#!/bin/bash
echo ">>>>>Ignore any errors.  Not all versions will build properly<<<<<\n"
rm -f ./sauced

echo "Building for Windows"
env GOOS=windows GOARCH=amd64 go build -ldflags "-X github.com/mdsauce/sauced/cmd.CurVersion=`git rev-parse --short HEAD`" -o sauced
mkdir -p ./builds/sauced-windows
mv sauced ./builds/sauced-windows
echo "done\n"

echo "Building for Linux"
env GOOS=linux GOARCH=amd64 go build -ldflags "-X github.com/mdsauce/sauced/cmd.CurVersion=`git rev-parse --short HEAD`" -o sauced
mkdir -p ./builds/sauced-linux
mv sauced ./builds/sauced-linux
echo "done\n"

echo "Building for Mac"
env GOOS=darwin GOARCH=amd64 go build -ldflags "-X github.com/mdsauce/sauced/cmd.CurVersion=`git rev-parse --short HEAD`" -o sauced
mkdir -p ./builds/sauced-mac
mv sauced ./builds/sauced-mac
echo "done\n"
