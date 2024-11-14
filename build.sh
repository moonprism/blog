#!/bin/sh

wget -O libsimple.zip https://github.com/wangfenjin/simple/releases/download/v0.4.0/libsimple-aarch64-linux-gnu-gcc-9.zip
unzip -j libsimple.zip -d ./libsimple

mkdir dict
wget -O dict/hmm_model.utf8 https://raw.githubusercontent.com/yanyiwu/cppjieba/refs/heads/master/dict/hmm_model.utf8
wget -O dict/idf.utf8 https://raw.githubusercontent.com/yanyiwu/cppjieba/refs/heads/master/dict/idf.utf8
wget -O dict/jieba.dict.utf8 https://raw.githubusercontent.com/yanyiwu/cppjieba/refs/heads/master/dict/jieba.dict.utf8
wget -O dict/stop_words.utf8 https://raw.githubusercontent.com/yanyiwu/cppjieba/refs/heads/master/dict/stop_words.utf8
wget -O dict/user.dict.utf8 https://raw.githubusercontent.com/yanyiwu/cppjieba/refs/heads/master/dict/user.dict.utf8

go build --tags fts5 -o blog
