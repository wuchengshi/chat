#!/bin/sh
rm -rf ./release
mkdir  release
go build -o chat
chmod +x ./chat
mv chat ./release/
cp favicon.ico ./release/
cp -arf ./asset ./release/
cp -arf ./view ./release/
cp -arf ./config ./release/