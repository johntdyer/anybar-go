#!/bin/bash
CGO_ENABLED=0
APPNAME=anybar-go

# Install dependencies
go get github.com/tools/godep

mkdir -p builds


# Build for OSX
godep go build -o builds/$APPNAME.osx
if [ $? -eq 0 ]; then
  echo "Success Build artifact - builds/$APPNAME.osx"
else
  echo "Build error"
  exit $?
fi

# Build for Linux
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 godep go build -o builds/$APPNAME.linux
if [ $? -eq 0 ]; then
  echo "Success Build artifact - builds/$APPNAME.linux"
else
  echo "Build error"
  exit $?
fi

chmod +x builds/$APPNAME.linux
chmod +x builds/$APPNAME.osx
