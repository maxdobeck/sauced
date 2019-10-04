#!/bin/bash
echo ">>>>>Ignore any errors.  Not all versions will build properly<<<<<\n"
rm -f ./sauced
mkdir -p ./builds/zips

echo "Building for Windows"
env GOOS=windows GOARCH=amd64 go build -ldflags "-X github.com/mdsauce/sauced/cmd.CurVersion=`git rev-parse --short HEAD`" -o sauced
mkdir -p ./builds/sauced-windows
mv sauced ./builds/sauced-windows
echo "done\n"

echo "Building for Linux"
env GOOS=linux GOARCH=amd64 go build -ldflags "-X github.com/mdsauce/sauced/cmd.CurVersion=`git rev-parse --short HEAD`" -o sauced
mkdir -p ./builds/sauced-linux
mv sauced ./builds/sauced-linux
tar cvzf sauced-linux.tar.gz  ./builds/sauced-linux/
mv sauced-linux.tar.gz ./builds/zips
echo "done\n"

echo "Building for Mac"
env GOOS=darwin GOARCH=amd64 go build -ldflags "-X github.com/mdsauce/sauced/cmd.CurVersion=`git rev-parse --short HEAD`" -o sauced
mkdir -p ./builds/sauced-mac
mv sauced ./builds/sauced-mac
tar cvzf sauced-mac.tar.gz  ./builds/sauced-mac/
mv sauced-mac.tar.gz ./builds/zips
echo "done\n"
