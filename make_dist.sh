#!/bin/sh

GOPATH=`pwd` go install github.com/fhermeni/wints
mkdir wints
cp bin/wints wints/
cp -r src/github.com/fhermeni/wints/static src/github.com/fhermeni/wints/wints.conf  wints/
tar cfz wints.tar.gz wints
rm -rf wints