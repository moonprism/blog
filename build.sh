#!/bin/sh

wget -O libsimple.zip https://github.com/wangfenjin/simple/releases/download/v0.4.0/libsimple-aarch64-linux-gnu-gcc-9.zip
unzip -j libsimple.zip -d ./dict

# go build --tags fts5 -o blog
